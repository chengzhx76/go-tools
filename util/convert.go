package util

import (
	"fmt"
	"math"
	"strings"
)

const (
	//dict10 = "0123456789"
	dict10 = "9876543210"
	dict16 = "0123456789ABCDEF"
	dict36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	dict38 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	dict64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

// hex 进制
func NumToHex(num int64, hex ...int32) string {
	h, dict := getHexDict(hex...)
	return fmt.Sprintf("%02d%s", h, HexConvert(num, dict)) // %02d 不够两位补0
}

func HexToNum(str string) int64 {
	hex := StringToInt32(SubString(str, 0, 2))
	_, dict := getHexDict(hex)
	return Convert10Hex(SubString(str, 2, 0), dict)
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

func getHexDict(hex ...int32) (int32, string) {
	h := getHex(hex...)
	dict := dict36
	if h == 10 {
		dict = dict10
	} else if h == 16 {
		dict = dict16
	} else if h == 36 {
		dict = dict36
	} else if h == 38 {
		dict = dict38
	} else if h == 64 {
		dict = dict64
	}
	return h, dict
}

func getHex(hex ...int32) int32 {
	if len(hex) > 0 {
		return hex[0]
	}
	return 36
}
