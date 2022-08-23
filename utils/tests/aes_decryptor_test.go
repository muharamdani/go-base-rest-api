//go:build unit || utils || aes || all

package tests

import (
	"github.com/go-playground/assert/v2"
	"testing"
)
import "github.com/muharamdani/go-base-rest-api/utils"

func TestAESDecrypt(t *testing.T) {
	key := "12345678901234567890123456789012"
	iv := "1234567890123456"
	ct := "aKTs8BBg4ZL2oiM2Hy/oPg=="
	res := utils.AESDecrypt(key, iv, ct)
	assert.Equal(t, "Test", res)
}
