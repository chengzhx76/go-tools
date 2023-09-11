package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(val string) string {
	h := md5.New()
	_, _ = io.WriteString(h, val)
	return fmt.Sprintf("%x", h.Sum(nil))
}
