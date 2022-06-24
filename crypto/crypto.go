package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

//20220616 leehs AES CTR 암호화
func Encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}
	mode := cipher.NewCTR(b, iv)
	mode.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}

//20220616 leehs AES CTR 복호화
func Decrypt(b cipher.Block, ciphertext []byte) []byte {
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCTR(b, iv)

	mode.XORKeyStream(plaintext, ciphertext)

	return plaintext
}
