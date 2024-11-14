package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
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

func Md5(val string) string {
	h := md5.New()
	_, _ = io.WriteString(h, val)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HmacSha1(data string, secret string) string {
	return HmacSha(sha1.New, data, secret)
}

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return HmacSha(sha256.New, data, secret)
}

func HmacSha(hashFunc func() hash.Hash, data string, secret string) string {
	h := hmac.New(hashFunc, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1HexDigest(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
