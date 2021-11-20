package util

import (
	"unicode"
	"unicode/utf8"
)

func ValiIdCard(id string) bool {
	//return idvalidator.IsValid(id, true)
	return true
}

func ValiChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func ValiChineseName(name string) bool {
	if ValiChineseChar(name) {
		return utf8.RuneCountInString(name) >= 2
	}
	return false
}

func ValiCarId(carId string, len int) bool {
	return utf8.RuneCountInString(TrimSpace(carId)) == len
}
