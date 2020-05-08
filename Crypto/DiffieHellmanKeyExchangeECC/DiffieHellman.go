package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

type ECCKey struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func (eccKey ECCKey) X509encodePrivateKey() []byte {
	bytes, _ := x509.MarshalECPrivateKey(eccKey.privateKey)
	return pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes})
}

func (eccKey *ECCKey) GenerateKeypair(curve elliptic.Curve, entropySource io.Reader) {
	var err error
	eccKey.privateKey, err = ecdsa.GenerateKey(curve, entropySource)
	if err != nil {
		panic(err.Error())
	}
	eccKey.publicKey = &eccKey.privateKey.PublicKey
	return
}

func (eccKey ECCKey) GenerateSharedSecret(publicKey *ecdsa.PublicKey) (sharedSecret [32]byte) {
	secret, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, eccKey.privateKey.D.Bytes())
	sharedSecret = sha256.Sum256(secret.Bytes())
	return
}

func main() {

	var selectedCurve elliptic.Curve = elliptic.P521()

	AliceKey := ECCKey{}
	AliceKey.GenerateKeypair(selectedCurve, rand.Reader)

	BobKey := ECCKey{}
	BobKey.GenerateKeypair(selectedCurve, rand.Reader)

	fmt.Printf("[DEBUG] Private key (Alice) %x \n", AliceKey.privateKey.D)
	fmt.Printf("[DEBUG] Private key (Bob)   %x \n", BobKey.privateKey.D)

	fmt.Printf("[DEBUG] Public key (Alice) (x:%x, y:%x) \n", AliceKey.publicKey.X, AliceKey.publicKey.Y)
	fmt.Printf("[DEBUG] Public key (Bob)   (x:%x, y:%x) \n", BobKey.publicKey.X, BobKey.publicKey.Y)

	secretAlice := AliceKey.GenerateSharedSecret(BobKey.publicKey)
	secretBob := BobKey.GenerateSharedSecret(AliceKey.publicKey)

	fmt.Printf("[INFO] Shared key (Alice) %x \n", secretAlice)
	fmt.Printf("[INFO] Shared key (Bob)   %x \n", secretBob)

}
