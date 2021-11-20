package util

import (
	"math"
	"strings"
)

const (
	dict38 = "0123456789abcdefghijklmnopqrstuvwxyz_-"
	dict64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

func From10To38(num int64) string {
	var str38 []byte
	for {
		var result byte
		var tmp []byte

		number := num % 38
		result  = dict38[number] // C

		// 临时变量，为了追加到头部
		tmp = append(tmp, result)

		str38 = append(tmp, str38...)
		num = num / 38

		if num == 0 {
			break
		}
	}
	return string(str38)
}

func From38To10(str38 string) int64 {
	var pos int
	var number int64
	str38Len := len(str38)

	for i := 0; i < str38Len; i++ {
		pos = strings.IndexAny(dict38, str38[i:i+1])
		number = int64(math.Pow(38,  float64(str38Len - i - 1)) * float64(pos)) + number
	}
	return number
}

func From10To64(num int64) string {
	var str64 []byte
	for {
		var result byte
		var tmp []byte

		number := num % 64 // 100%64 = 64
		result  = dict64[number] // C

		// 临时变量，为了追加到头部
		tmp = append(tmp, result)

		str64 = append(tmp, str64...)
		num = num / 64

		if num == 0 {
			break
		}
	}
	return string(str64)
}

func From64To10(str64 string) int64 {
	var pos int
	var number int64
	str64Len := len(str64)

	for i := 0; i < str64Len; i++ {
		pos = strings.IndexAny(dict64, str64[i:i+1])
		number = int64(math.Pow(64,  float64(str64Len - i - 1)) * float64(pos)) + number
	}
	return number
}