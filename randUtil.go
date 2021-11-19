package tool

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 指定范围的随机数
func RangeInt(min int, max int) int {
	return rand.Intn(max - min) + min
}