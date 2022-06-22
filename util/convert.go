package util

import (
	"math"
	"strings"
)

const (
	//dict10 = "0123456789"
	dict10 = "9876543210"
	dict38 = "0123456789abcdefghijklmnopqrstuvwxyz_-"
	dict64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

func From10To10Str(num int64) string {
	return HexConvert(num, dict10)
}

func From10StrTo10(str10 string) int64 {
	return Convert10Hex(str10, dict10)
}

func From10To38(num int64) string {
	return HexConvert(num, dict38)
}

func From38To10(str38 string) int64 {
	return Convert10Hex(str38, dict38)
}

func From10To64(num int64) string {
	return HexConvert(num, dict64)
}

func From64To10(str64 string) int64 {
	return Convert10Hex(str64, dict64)
}

/*
	tenHexNum: 10进制数
	toHex: 要转换的进制数
	dict: 要转换的进制数的字典
*/
func HexConvert(tenHexNum int64, dict string) string {
	var str []byte
	toHex := int64(len(dict))
	for {
		var result byte
		var tmp []byte

		number := tenHexNum % toHex
		result = dict[number] // C

		// 临时变量，为了追加到头部
		tmp = append(tmp, result)

		str = append(tmp, str...)
		tenHexNum = tenHexNum / toHex

		if tenHexNum == 0 {
			break
		}
	}
	return string(str)
}

/*
	转换成 10 进制
	str: 要转换的字符串
	dict: 要转换的进制数的字典
*/
func Convert10Hex(str, dict string) int64 {
	var pos int
	var number int64
	originalHex := len(dict) // 原本进制数
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		pos = strings.IndexAny(dict, str[i:i+1])
		number = int64(math.Pow(float64(originalHex), float64(strLen-i-1))*float64(pos)) + number
	}
	return number
}
