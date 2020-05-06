package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func SHA256SUM(key string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

func encrypt(data []byte, encryptionKey []byte) []byte {
	block, _ := aes.NewCipher(encryptionKey)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, decryptionKey []byte) []byte {
	block, err := aes.NewCipher(decryptionKey)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, encryptionKey []byte, outFileName string) {
	data, _ := ioutil.ReadFile(filename)
	encryptedData := encrypt(data, encryptionKey)
	f, _ := os.Create(outFileName)
	defer f.Close()
	f.Write([]byte(b64.StdEncoding.EncodeToString(encryptedData)))
}

func decryptFile(encryptedFilename string, decryptionKey []byte, decryptedFileName string) {
	data, _ := ioutil.ReadFile(encryptedFilename)
	encryptedData, _ := b64.StdEncoding.DecodeString(string(data))
	decryptedData := decrypt(encryptedData, decryptionKey)
	f, _ := os.Create(decryptedFileName)
	defer f.Close()
	f.Write(decryptedData)
}

func localTest() {

	text := "Alice Sends to Bob a really long message to test if the encryption works on long messages and to test the limits of bob's tolerance "
	passphrase := "Hello, World!"

	encryptionKey := SHA256SUM(passphrase)
	fmt.Println("[DEBUG] Encryption Key Used : ", hex.EncodeToString(encryptionKey))
	ciphertext := encrypt([]byte(text), encryptionKey)
	fmt.Println("[DEBUG] Ciphertext : ", hex.EncodeToString(ciphertext))
	plaintext := string(decrypt(ciphertext, encryptionKey))
	fmt.Println("[DEBUG] Plaintext : ", plaintext)

	if text == plaintext {
		fmt.Println("[INFO] Local Test Passed")
	} else {
		panic("[ERROR] Local Test Failed")
	}

}

func fileTest() {

	encryptionKey := SHA256SUM("Passphrase")
	encryptFile("Crypto/SymmetricKey/TestFile.txt", encryptionKey, "Crypto/SymmetricKey/enc-TestFile.txt")
	decryptFile("Crypto/SymmetricKey/enc-TestFile.txt", encryptionKey, "Crypto/SymmetricKey/dec-TestFile.txt")

	inputFile, _ := ioutil.ReadFile("Crypto/SymmetricKey/TestFile.txt")
	inputDigest := hex.EncodeToString(SHA256SUM(string(inputFile)))

	outputFile, _ := ioutil.ReadFile("Crypto/SymmetricKey/dec-TestFile.txt")
	outputDigest := hex.EncodeToString(SHA256SUM(string(outputFile)))

	fmt.Println("[DEBUG] Input File Digest  : ", inputDigest)
	fmt.Println("[DEBUG] Output File Digest : ", outputDigest)

	if inputDigest == outputDigest {
		fmt.Println("[INFO] File Test Passed")
	} else {
		panic("[ERROR] File Test Failed")
	}

}

func main() {

	localTest()
	fileTest()

}
