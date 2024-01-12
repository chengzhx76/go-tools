package util

import (
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"regexp"
	"sort"
)

func Contains(coll []string, value string) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return IndexOf(coll, value) >= 0
	}
	return false
}

func RegexContains(coll []string, value string) bool {
	if Contains(coll, value) {
		return true
	} else {
		for _, collRegx := range coll {
			regex := regexp.MustCompile(collRegx)
			if regex.MatchString(value) {
				return true
			}
		}
	}
	return false
}

func Uint8Contains(coll []uint8, value uint8) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return Uint8IndexOf(coll, value) >= 0
	}
	return false
}

func Int64Contains(coll []int64, value int64) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return Int64IndexOf(coll, value) >= 0
	}
	return false
}

func BoolContains(coll []bool, value bool) bool {
	if !IsNil(coll) && len(coll) > 0 {
		return BoolIndexOf(coll, value) >= 0
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
func Int64IndexOf(coll []int64, value int64) int {
	for index, item := range coll {
		if item == value {
			return index
		}
	}
	return -1
}
func BoolIndexOf(coll []bool, value bool) int {
	for index, item := range coll {
		if item == value {
			return index
		}
	}
	return -1
}

// 移除集合中的值, 返回移除之后的数组
func CollectionRemove(coll []string, value string) []string {
	if index := IndexOf(coll, value); index > -1 {
		return append(coll[:index], coll[index+1:]...)
	}
	return coll
}

// 移除集合中的值, 返回移除之后的数组
func RemoveStrElms(sl []string, elms ...string) []string {
	if len(sl) == 0 || len(elms) == 0 {
		return sl
	}
	// 先将元素转为 set
	m := make(map[string]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 过滤掉指定元素
	res := make([]string, 0, len(sl))
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

// 移除指定位置的元素，返回新切片
func SliceRemoveInt64(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}

func SliceRemove(s []any, index int) []any {
	return append(s[:index], s[index+1:]...)
}

func SortInt64(s []int64) []int64 {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

// 数组已符号链接
// Deprecated 建议使用 strings.Join
func CollectionSymbolJoin(coll []string, symbol string) string {
	result := SYMBOL_EMPTY
	for i, item := range coll {
		if !IsBlank(item) {
			if i == 0 {
				result += item
				continue
			}
			result += fmt.Sprintf("%s%s", symbol, item)
		}
	}
	return result
}

// 数组已符号链接
func Uint8CollectionSymbolJoin(coll []uint8, symbol string) string {
	result := SYMBOL_EMPTY
	for i, item := range coll {
		if !IsNil(item) {
			if i == 0 {
				result += Uint8ToString(item)
				continue
			}
			result += fmt.Sprintf("%s%d", symbol, item)
		}
	}
	return result
}

// 数组已符号链接
func IntCollectionSymbolJoin(coll []int, symbol string) string {
	result := SYMBOL_EMPTY
	for i, item := range coll {
		if !IsNil(item) {
			if i == 0 {
				result += IntToString(item)
				continue
			}
			result += fmt.Sprintf("%s%d", symbol, item)
		}
	}
	return result
}
