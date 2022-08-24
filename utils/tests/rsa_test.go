//go:build unit || utils || rsa || all

package tests

import (
	"github.com/muharamdani/go-base-rest-api/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var msg = "secret"
var encMsg string
var cwd string

func init() {
	dir, _ := os.Getwd()
	cwd = dir
}

func TestGenerateKeyPairs(t *testing.T) {
	// Generate public and private key
	utils.GenerateKeyPairs()
	// Check if public key are generated
	_, errPub := os.ReadFile("./public.pem")
	if errPub != nil {
		assert.Error(t, errPub, "Public file not found")
	}
	// Check if private key are generated
	_, errPri := os.ReadFile("./private.pem")
	if errPri != nil {
		assert.Error(t, errPri, "Private file not found")
	}
	assert.True(t, true, "Key pair generated")
}

func TestRSAEncrypt(t *testing.T) {
	pubKeyLoc := cwd + "/public.pem"
	pubKey := utils.ByteToPublicKey(pubKeyLoc)
	encMsg = utils.RSAEncrypt(msg, pubKey)
	assert.NotEmptyf(t, encMsg, "Encrypted message generated!")
	if err := os.Remove("./public.pem"); err != nil {
		t.Errorf("Error removing file: %s", err)
	}
}

func TestRSADecrypt(t *testing.T) {
	privKeyLoc := cwd + "/private.pem"
	privKey := utils.BytesToPrivateKey(privKeyLoc)
	res := utils.RSADecrypt(encMsg, privKey)
	assert.Equal(t, msg, res, "Decrypted message is equal to original message")
	if err := os.Remove("./private.pem"); err != nil {
		t.Errorf("Error removing file: %s", err)
	}
}
