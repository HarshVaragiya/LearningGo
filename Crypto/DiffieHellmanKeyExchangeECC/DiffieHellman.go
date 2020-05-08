package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func GenerateKeypair(curve elliptic.Curve, entropySource io.Reader) (privateKey *ecdsa.PrivateKey, publicKey ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(curve, entropySource)
	if err != nil {
		panic(err.Error())
	}
	publicKey = privateKey.PublicKey
	return
}

func GenerateSharedSecret(privatekey *ecdsa.PrivateKey, publicKey ecdsa.PublicKey) (sharedSecret [32]byte) {
	secret, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privatekey.D.Bytes())
	sharedSecret = sha256.Sum256(secret.Bytes())
	return
}

func main() {
	var selectedCurve elliptic.Curve = elliptic.P521()

	privateKeyAlice, publicKeyAlice := GenerateKeypair(selectedCurve, rand.Reader)
	privateKeyBob, publicKeyBob := GenerateKeypair(selectedCurve, rand.Reader)

	fmt.Printf("[DEBUG] Private key (Alice) %x \n", privateKeyAlice.D)
	fmt.Printf("[DEBUG] Private key (Bob)   %x \n", privateKeyBob.D)

	fmt.Printf("[DEBUG] Public key (Alice) (x:%x, y:%x) \n", publicKeyAlice.X, publicKeyAlice.Y)
	fmt.Printf("[DEBUG] Public key (Bob)   (x:%x, y:%x) \n", publicKeyBob.X, publicKeyBob.Y)

	secretAlice := GenerateSharedSecret(privateKeyAlice, publicKeyBob)
	secretBob := GenerateSharedSecret(privateKeyBob, publicKeyAlice)

	fmt.Printf("[INFO] Shared key (Alice) %x \n", secretAlice)
	fmt.Printf("[INFO] Shared key (Bob)   %x \n", secretBob)

}
