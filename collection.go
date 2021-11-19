package tool

import (
	"fmt"
)

func Contains(coll []string, value string) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return IndexOf(coll, value) >= 0
	}
	return false
}

func Uint8Contains(coll []uint8, value uint8) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return Uint8IndexOf(coll, value) >= 0
	}
	return false
}

func IndexOf(coll []string, value string) int {
	for index, item := range coll {
		if item == value {
			return index
		}
	}
	return -1
}

func Uint8IndexOf(coll []uint8, value uint8) int {
	for index, item := range coll {
		if item == value {
			return index
		}
	}
	return -1
}

// 移除集合中的值
func CollectionRemove(coll []string, value string) []string {
	if index := IndexOf(coll, value); index > -1 {
		return append(coll[:index], coll[index+1:]...)
	}
	return coll
}

func SliceRemove(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}

// 数组已符号链接
func CollectionSymbolJoin(coll []string, symbol string) string {
	result := ""
	for _, item := range coll {
		if !IsBlank(item) {
			result += fmt.Sprintf("%v%v", item, symbol)
		}
	}
	return result
}

// 数组已符号链接
func Uint8CollectionSymbolJoin(coll []uint8, symbol string) string {
	result := ""
	for _, item := range coll {
		if !IsNil(item) {
			result += fmt.Sprintf("%d%s", item, symbol)
		}
	}
	return result
}

// 数组已符号链接
func IntCollectionSymbolJoin(coll []int, symbol string) string {
	result := ""
	for _, item := range coll {
		if !IsNil(item) {
			result += fmt.Sprintf("%d%s", item, symbol)
		}
	}
	return result
}
