package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	secretMessage := []byte("This message is secret")
	label := []byte("orders")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating RSA key: %s\n", err)
	}

	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	fmt.Printf("CipherText: %x\n", cipherText)
}
