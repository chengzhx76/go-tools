package tests

import (
	"fmt"
	"github.com/chengzhx76/go-tools/util"
	"testing"
)

func Test_DES(t *testing.T) {

	plaintext := "cheng"
	//ciphertext := "7KKkW7fofd3HF3M+J5vkRQ=="
	key := "chengzhx76"

	ciphertext, err := util.DESEncrypter(plaintext, key)
	decryptedText, err2 := util.DESDecryption(ciphertext, key)

	fmt.Printf("err: %s\n", err)
	fmt.Printf("err2: %s\n", err2)
	fmt.Printf("ciphertext: %s\n", ciphertext)
	fmt.Printf("decryptedText: %s\n", decryptedText)

}
