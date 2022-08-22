//go:build unit || utils || aes || all

package tests

import "testing"
import "github.com/muharamdani/go-base-rest-api/utils"


func TestAESDecrypt(t *testing.T) {
	key := []byte("1234567890123456")
	ct := "ciphertext"
	utils.AESDecrypt(key, ct)
}
