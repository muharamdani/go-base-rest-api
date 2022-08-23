package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func AESDecrypt(key, iv, ct string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(ct)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	fmt.Println(string(ciphertext))
	return string(ciphertext)
}
