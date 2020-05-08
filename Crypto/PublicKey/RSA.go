package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

func keygen(keySize int) (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err.Error())
	}
	publicKey = &privateKey.PublicKey
	return
}

func x509encodePrivateKey(key *rsa.PrivateKey) (bytes []byte) {
	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	bytes = pem.EncodeToMemory(pemPrivateBlock)
	return
}
func x509encodePublicKey(key *rsa.PublicKey) (bytes []byte) {
	var pemPublicBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key),
	}
	bytes = pem.EncodeToMemory(pemPublicBlock)
	return
}

func encrypt(message []byte, key *rsa.PublicKey) (ciphertext []byte) {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, key, message, label)
	if err != nil {
		panic(err.Error())
	}
	return
}

func sign(message []byte, key *rsa.PrivateKey) (signature []byte) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	hasher := crypto.SHA256
	pssh := hasher.New()
	pssh.Write(message)
	checksum := pssh.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, key, hasher, checksum, &opts)
	if err != nil {
		panic(err.Error())
	}
	return
}

func decrypt(ciphertext []byte, key *rsa.PrivateKey) (plaintext []byte) {
	label := []byte("")
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, key, ciphertext, label)
	if err != nil {
		panic(err.Error())
	}
	return
}

func verify(message []byte, key *rsa.PublicKey, signature []byte) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	hasher := crypto.SHA256
	pssh := hasher.New()
	pssh.Write(message)
	checksum := pssh.Sum(nil)
	err := rsa.VerifyPSS(key, hasher, checksum, signature, &opts)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("[INFO] Signature Verified")
	}
}

func main() {
	message := "Hello, World!"
	privateKeyAlice, publicKeyAlice := keygen(2048)
	fmt.Println("[DEBUG] Keys  generated by Alice  : ")
	fmt.Println(string(x509encodePrivateKey(privateKeyAlice)))
	fmt.Println(string(x509encodePublicKey(publicKeyAlice)))

	privateKeyBob, publicKeyBob := keygen(2048)
	fmt.Println("[DEBUG] Keys generated by Bob : ")
	fmt.Println(string(x509encodePrivateKey(privateKeyBob)))
	fmt.Println(string(x509encodePublicKey(publicKeyBob)))

	fmt.Println("[INFO] Bob Encrypts the Text Message with Alice's public key and signs with Bob's private key")
	ciphertext := encrypt([]byte(message), publicKeyAlice)
	signature := sign([]byte(message), privateKeyBob)

	// Mess with this to test signature verification
	signature, _ = hex.DecodeString(hex.EncodeToString(signature)[:])

	fmt.Println("[DEBUG] Ciphertext : ", hex.EncodeToString(ciphertext))
	fmt.Println("[DEBUG] Signature  : ", hex.EncodeToString(signature))

	fmt.Println("[INFO] Alice Decrypts the ciphertext with her private key and verifies signature with Bob's public key")
	plaintext := decrypt(ciphertext, privateKeyAlice)
	fmt.Println("[INFO] Plaintext decrypted : ", string(plaintext))
	verify(plaintext, publicKeyBob, signature)

}
