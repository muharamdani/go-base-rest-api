package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

// GenerateKeyPairs generate public and private key
func GenerateKeyPairs() {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA keyn")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s n", err)
		os.Exit(1)
	}
}

// RSAEncrypt Encrypt the data with the public key
func RSAEncrypt(secretMessage string, key rsa.PublicKey) string {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), nil)
	if err != nil {
		fmt.Printf("error when encrypting: %s n", err)
		os.Exit(1)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// RSADecrypt Decrypt the data with the private key
func RSADecrypt(cipherText string, privKey rsa.PrivateKey) string {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	plaintext, err := privKey.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		fmt.Printf("error when decrypting: %s n", err)
		os.Exit(1)
	}
	return string(plaintext)
}

func BytesToPrivateKey(keyLocation string) rsa.PrivateKey {
	// Default placement private.pem is in root directory
	if keyLocation == "" {
		rootPath := GetRootPath()
		keyLocation = rootPath + "private.pem"
	}

	priKey, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(priKey)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return *key
}

func ByteToPublicKey(keyLocation string) rsa.PublicKey {
	// Default placement public.pem is in root directory
	if keyLocation == "" {
		rootPath := GetRootPath()
		keyLocation = rootPath + "public.pem"
	}

	pubKey, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(pubKey)
	key, _ := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return *key.(*rsa.PublicKey)
}
