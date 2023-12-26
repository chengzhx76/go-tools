package util

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

var HASH_CHECK_FAIL = errors.New("check hash err")

// sha256 校验
func Sha1HashCheck(bodyBytes []byte, hash string) error {
	h := sha1.New()
	h.Write(bodyBytes)
	sha256Val := fmt.Sprintf("%x", h.Sum(nil))
	if sha256Val != hash {
		fmt.Println("check hash fail")
		return HASH_CHECK_FAIL
	}
	return nil
}
