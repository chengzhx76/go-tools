package util

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	. "github.com/chengzhx76/go-tools/consts"
)

const iv = "12345678"

// 加密
func DESEncrypter(plaintext, key string) (string, error) {
	plaintextBytes := []byte(plaintext)
	key = SubString(key, 0, 8)
	keyBytes := []byte(key)
	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return SYMBOL_EMPTY, err
	}
	bs := block.BlockSize()
	plaintextBytes = PKCS5Padding(plaintextBytes, bs)
	if len(plaintextBytes)%bs != 0 {
		return SYMBOL_EMPTY, errors.New("need a multiple of the blocksize")
	}
	ciphertext := make([]byte, len(plaintextBytes))
	dst := ciphertext
	for len(plaintextBytes) > 0 {
		block.Encrypt(dst, plaintextBytes[:bs])
		plaintextBytes = plaintextBytes[bs:]
		dst = dst[bs:]
	}

	return Base64Encode(ciphertext), nil
}

// PKCS5Padding 对明文进行填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)

}

// 解密
func DESDecryption(cipherText, key string) (string, error) {
	key = SubString(key, 0, 8)
	keyBytes := []byte(key)

	cipherTextBytes, err := Base64Decode(cipherText)
	if err != nil {
		return SYMBOL_EMPTY, err
	}
	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return SYMBOL_EMPTY, err
	}
	decryptedText := make([]byte, len(cipherTextBytes))
	dst := decryptedText
	bs := block.BlockSize()
	if len(cipherTextBytes)%bs != 0 {
		return SYMBOL_EMPTY, errors.New("crypto/cipher: input not full blocks")
	}
	for len(cipherTextBytes) > 0 {
		block.Decrypt(dst, cipherTextBytes[:bs])
		cipherTextBytes = cipherTextBytes[bs:]
		dst = dst[bs:]
	}
	decryptedText = PKCS5UnPadding(decryptedText)

	return string(decryptedText), nil
}

func PKCS5UnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})

}

// Base64 解码
func Base64Decode(encodeStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodeStr)
}

// Base64 编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
