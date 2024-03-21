package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

//https://www.cnblogs.com/lpgit/p/10632756.html

func IsBlank(str string) bool {
	str = TrimSpace(str)
	return len(str) == 0 || str == SYMBOL_EMPTY
}

// 是否全部为空，全部为空返回 true,有一个不为空返回 false
func IsAllBlank(strs ...string) bool {
	flags := make([]bool, len(strs))
	for i, v := range strs {
		flags[i] = IsBlank(v)
	}
	return !BoolContains(flags, false)
}

// 是否全部不为空，全部不为空返回 true, 有一个为空返回 false
func IsAllNotBlank(strs ...string) bool {
	flags := make([]bool, len(strs))
	for i, v := range strs {
		flags[i] = !IsBlank(v)
	}
	return !BoolContains(flags, false)
}

func IsHasBlank(strs ...string) bool {
	for _, v := range strs {
		if IsBlank(v) {
			return true
		}
	}
	return false
}

// 有一个不等于空的返回true(只要有值就返回true)
func IsHasNotBlank(strs ...string) bool {
	for _, v := range strs {
		if !IsBlank(v) {
			return true
		}
	}
	return false
}

func RandString(length ...int) string {
	num := 32
	if len(length) > 0 {
		num = length[0]
	}
	if num < 10 {
		num = 10
	}
	return GenerateBytesUUID(num)
}

func GenerateBytesUUID(len int) string {
	uuid := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		panic(fmt.Sprintf("USER: Error generating UUID: %s", err))
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	uuidString := idBytesToStr(uuid, len)
	return uuidString
}
func idBytesToStr(id []byte, length int) string {
	str := fmt.Sprintf("%x%x%x%x%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
	return str[:length]
}

func NilToBlank(data any) string {
	if IsNil(data) {
		return SYMBOL_EMPTY
	}
	val, ok := data.(string)
	if ok {
		return val
	} else {
		log.Println("data not string type")
	}
	return SYMBOL_EMPTY
}

// 替换 url 中的占位符
func Replace(s string, value map[string]string) string {
	newString := s
	for k, v := range value {
		newString = strings.Replace(newString, k, v, 1)
	}
	return newString
}

func ParamEncode(params map[string]string, start string) string {
	var args = start
	for key, value := range params {
		args += key + SYMBOL_EQUAL + value + SYMBOL_AND
	}
	return TrimSuffix(args, SYMBOL_AND)
}

// 去掉 suffix
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// 去掉所有空格
func TrimSpace(str string) string {
	return strings.Replace(str, SYMBOL_SPACE, SYMBOL_EMPTY, -1)
}

// 删除指定位置元素
func RemoveElement(elems []string, index int32) []string {
	return append(elems[:index], elems[index+1:]...)
}

// s 根据 sep 拆分后 获取第 index(从后边开始数) 个元素 从 0 开始
func SplitSuffix(s, sep string, index int) string {
	vals := strings.Split(s, sep)
	if index > len(vals)-1 {
		return SYMBOL_EMPTY
	}
	return vals[len(vals)-1-index]
}

// s 根据 sep 拆分后 获取第 index 个元素 从 0 开始
func Split(s, sep string, index int) string {
	vals := strings.Split(s, sep)
	if index > len(vals)-1 {
		return SYMBOL_EMPTY
	}
	return vals[index]
}

// 查找字符串位置
func StringIndexOf(str, substr string) int {
	// 子串在字符串的字节位置
	index := strings.Index(str, substr)
	if index >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:index]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		index = len(rs)
	}

	return index
}

// 包含字符串
func StrContains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// 从第一位开始截取，返回截取的字符串
// length 截取的个数
// 废弃 建议使用 SubString
// Deprecated
func SubBeforeString(s string, length int) string {
	endIndex := length
	if endIndex > len(s) {
		endIndex = len(s)
	}
	return s[:endIndex]
}

// 从最后一位开始截取，返回截取的字符串
// length 截取的个数
// 废弃 建议使用 SubString
// Deprecated
func SubAfterString(s string, length int) string {

	endIndex := len(s)
	startIndex := endIndex - length
	if startIndex < 0 {
		startIndex = 0
	}

	return s[startIndex:endIndex]
}

// start：起始下标，负数从尾部开始。这里的-表示反向读取，从1开始；正数从0开始
// length：截取长度，负数表示反向截取；0 表示截取到末尾
// SubString("0123456789", 4, 2) //  = 45 // 从第4索引个开始截取2个
// SubString("0123456789", 2, -1) //  = 1
// SubString("0123456789", -1, -1) // = 8
// SubString("0123456789", -3, 3) // = 789
// SubString("0123456789", -1, 3) // = 9
// SubString("0123456789", -3, -3) // = 456
// SubString("0123456789", 0, -3) // =
// SubString("0123456789", -3, 0) // = 789
// SubString2("0123456789", 0, 0) // = 0123456789
func SubString(str string, start, length int) string {
	s := []rune(str)
	totalLen := len(s)
	if totalLen == 0 {
		return SYMBOL_EMPTY
	}
	startIndex := start
	endIndex := totalLen
	// 允许从尾部开始计算
	if start > 0 {
		if length > 0 {
			endIndex = startIndex + length
		} else if length < 0 {
			startIndex = startIndex - (-length)         // 绝对值
			endIndex = startIndex + (start - (-length)) // 绝对值
		}
	} else if start < 0 {
		if length > 0 {
			startIndex = endIndex - (-start) // 绝对值
			endIndex = startIndex + length
		} else if length < 0 {
			startIndex = endIndex - (-start) - (-length) // 绝对值
			endIndex = startIndex + (-length)            // 绝对值
		} else if length == 0 {
			startIndex = totalLen - (-start) // 绝对值
			endIndex = totalLen
		}
	} else if start == 0 {
		if length > 0 {
			endIndex = startIndex + length
		} else if length < 0 || length == 0 {
			return SYMBOL_EMPTY
		}
	}
	if startIndex > totalLen {
		return SYMBOL_EMPTY
	}
	if endIndex > totalLen {
		return string(s[startIndex:])
	} else {
		return string(s[startIndex:endIndex])
	}
}

// start：起始下标，负数从尾部开始，-1为最后一个
// length：截取长度，负数表示截取到末尾
// ex: util.SubString("0123456789", 4, 2) = 45
// ex: util.SubString("0123456789", 2, -1) = 23456789
// ex: util.SubString("0123456789", -1, -1) = 9
// ex: util.SubString("0123456789", -3, 3) = 789
/*func SubString(str string, start, length int) string {
	s := []rune(str)
	totalLen := len(s)
	if totalLen == 0 {
		return SYMBOL_EMPTY
	}
	startIndex := start
	// 允许从尾部开始计算
	if start < 0 {
		startIndex = totalLen + start
		if startIndex < 0 {
			return SYMBOL_EMPTY
		}
	}
	if startIndex > totalLen {
		return SYMBOL_EMPTY
	}
	endIndex := startIndex + length
	if length < 0 {
		endIndex = totalLen
	}
	if endIndex > totalLen {
		return string(s[startIndex:])
	} else {
		return string(s[startIndex:endIndex])
	}
}*/

// 翻转切片 [8 6 7 5 3 0 9] reversed: [9 0 3 5 7 6 8]
func ReverseStrings(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(ReverseStrings(input[1:]), input[0])
}

func JSONMarshal(v any, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte(SYMBOL_LT), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(SYMBOL_GT), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte(SYMBOL_AND), -1)
	}
	return b, err
}

func ValidateValNotNil(body map[string]any, key string) error {
	valObj := body[key]
	if IsNil(valObj) {
		return errors.New(fmt.Sprintf("<%s> value is nil"))
	}
	return nil
}

func ValStrNotNil(body map[string]any, key string) (string, error) {
	err := ValidateValNotNil(body, key)
	return ValStr(body, key), err
}

func ValStr(body map[string]any, key string, def ...string) string {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return SYMBOL_EMPTY
	}
	return AnyToString(body[key], def...)
}

func ValString(body map[string]string, key string, def ...string) string {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return SYMBOL_EMPTY
	}
	return AnyToString(body[key], def...)
}

func ValSlice(body map[string]any, key string, def ...[]any) []any {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return nil
	}
	return AnyToSlice(body[key], def...)
}

func ValFloat64(body map[string]any, key string, def ...float64) float64 {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	return AnyToFloat64(valObj)
}

// 大整数精度丢失
func ValJsonNumberToInt64(body map[string]any, key string, def ...int64) int64 {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	val, ok := valObj.(json.Number)
	if ok {
		i64, err := val.Int64()
		if err != nil {
			log.Printf(fmt.Sprintf("json.Number to Int64 err %s", err.Error()))
		}
		return i64
	} else {
		if len(def) > 0 {
			return def[0]
		}
		keyType, keyValue := reflect.TypeOf(valObj), reflect.ValueOf(valObj)
		log.Printf("<%s> is <%v> not json.Number type return default val 0 value<%v>", key, keyType, keyValue)
	}
	return 0
}

// 大整数精度丢失
func ValJsonNumberToString(body map[string]any, key string, def ...string) string {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return SYMBOL_EMPTY
	}
	valObj := body[key]
	val, ok := valObj.(json.Number)
	if ok {
		return val.String()
	} else {
		if len(def) > 0 {
			return def[0]
		}
		keyType, keyValue := reflect.TypeOf(valObj), reflect.ValueOf(valObj)
		log.Printf("<%s> is <%v> not json.Number type return default val '' value<%v>", key, keyType, keyValue)
	}
	return SYMBOL_EMPTY
}

// Deprecated ValInt32
func ValFloat64ToInt32(body map[string]any, key string, def ...int32) int32 {
	return ValInt32(body, key, def...)
}

// Deprecated ValInt64
func ValFloat64ToInt64(body map[string]any, key string, def ...int64) int64 {
	return ValInt64(body, key, def...)
}

func ValUnit8(body map[string]any, key string, def ...uint8) uint8 {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	return AnyToUint8(valObj, def...)
}

func ValInt(body map[string]any, key string, def ...int) int {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	return AnyToInt(valObj, def...)
}

func ValInt32(body map[string]any, key string, def ...int32) int32 {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	return AnyToInt32(valObj, def...)
}

func ValInt64(body map[string]any, key string, def ...int64) int64 {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}
	valObj := body[key]
	return AnyToInt64(valObj, def...)
}

func ValBool(body map[string]any, key string, def ...bool) bool {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return false
	}
	valObj := body[key]
	return AnyToBool(valObj, def...)
}

func ValMap(body map[string]any, key string, def ...map[string]any) map[string]any {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return nil
	}
	valObj := body[key]
	return AnyToMap(valObj, def...)
}

func ValAny(body map[string]any, key string, def ...any) any {
	if body == nil {
		if len(def) > 0 {
			return def[0]
		}
		return nil
	}
	valObj := body[key]
	return valObj
}

func AnyToMap(data any, defVal ...map[string]any) map[string]any {
	val, ok := data.(map[string]any)
	if ok {
		return val
	} else {
		if len(defVal) > 0 && defVal[0] != nil {
			return defVal[0]
		}
		keyType := reflect.TypeOf(data)
		log.Printf("InterfaceToMap data is<%v> not map type return default val nil map", keyType)
	}
	return val
}

func AnyToString(data any, defVal ...string) string {
	val := NilToBlank(data)
	if !IsBlank(val) {
		return val
	} else {
		if len(defVal) > 0 {
			return defVal[0]
		}
	}
	return val
}

func AnyToUint8(data any, defVal ...uint8) uint8 {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		log.Println("AnyToUint8 data is nil ret default 0")
		return 0
	}
	val, ok := data.(float64)
	if ok {
		return uint8(val)
	} else {
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToUint8 data is <%v> not num type ret default 0", keyType)
		return 0
	}
}

func AnyToInt(data any, defVal ...int) int {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		log.Println("AnyToInt data is nil ret default 0")
		return 0
	}
	val, ok := data.(float64)
	if ok {
		return int(val)
	} else {
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToInt data is <%v> not num type ret default 0", keyType)
		return 0
	}
}

func AnyToInt32(data any, defVal ...int32) int32 {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		log.Println("InterfaceToInt data is nil ret default 0")
		return 0
	}
	val, ok := data.(float64)
	if ok {
		return int32(val)
	} else {
		keyType := reflect.TypeOf(data)
		log.Printf("InterfaceToInt data is <%v> not num type ret default 0", keyType)
		return 0
	}
}

func AnyToInt64(data any, defVal ...int64) int64 {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToInt64 data is <%v> not num type ret default 0", keyType)
		return 0
	}

	val := int64(0)

	switch data.(type) {
	case float64:
		val = int64(data.(float64))
	case string:
		val = StringToInt64(data.(string))
	default:
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToInt64 data is <%v> not num type ret default 0", keyType)
	}
	return val
}

func AnyToFloat64(data any, defVal ...float64) float64 {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToFloat64 data is <%v> not num type ret default 0", keyType)
		return 0
	}

	val := float64(0)

	switch data.(type) {
	case float64:
		val = float64(data.(float64))
	case string:
		val = StringToFloat64(data.(string))
	default:
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToInt64 data is <%v> not num type ret default 0", keyType)
	}
	return val
}

func AnyToBool(data any, defVal ...bool) bool {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToBool data is <%v> not bool type ret default false", keyType)
		return false
	}

	val := false
	switch data.(type) {
	case bool:
		val = data.(bool)
	default:
		keyType := reflect.TypeOf(data)
		log.Printf("AnyToBool data is <%v> not bool type ret default false", keyType)
	}
	return val
}

func AnyToSlice(data any, defVal ...[]any) []any {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
	}
	val, ok := data.([]any)
	if ok {
		return val
	} else {
		if len(defVal) > 0 {
			return defVal[0]
		}
		keyType, keyValue := reflect.TypeOf(data), reflect.ValueOf(data)
		log.Printf("[error] <%v> not slice type return default val value<%v>", keyType, keyValue)
	}
	return nil
}

// 废弃 建议使用 AnyToMap
// Deprecated
func InterfaceToMap(data any, defVal ...map[string]any) map[string]any {
	return AnyToMap(data, defVal...)
}

// 废弃 建议使用 AnyToString
// Deprecated
func InterfaceToString(data any) string {
	return AnyToString(data)
}

// 废弃 建议使用 AnyToInt
// Deprecated
func InterfaceToInt(data any, defVal ...int) int {
	return AnyToInt(data, defVal...)
}

// 废弃 建议使用 AnyToInt64
// Deprecated
func InterfaceToInt64(data any, defVal ...int64) int64 {
	return AnyToInt64(data, defVal...)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
func Int32ToString(i int32) string {
	return strconv.Itoa(int(i))
}

func Uint8ToString(i uint8) string {
	return strconv.Itoa(int(i))
}

func BoolToString(boolVal bool) string {
	return strconv.FormatBool(boolVal)
}

func StringToBool(boolStr string) bool {
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		log.Fatal(fmt.Sprintf("string to bool err<%s>", err.Error()))
	}
	return boolVal
}

func Uint8SliceToStringSlice(is []uint8) []string {
	ss := make([]string, len(is))
	for i, v := range is {
		ss[i] = Uint8ToString(v)
	}
	return ss
}

// https://www.cnblogs.com/f-ck-need-u/p/9863915.html

func StringToUint8(s string) uint8 {
	if IsBlank(s) {
		log.Println("string s is nil return 0")
		return UNKNOWN
	}
	u64, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		log.Println("string to uint8 error", err)
		return UNKNOWN
	}
	u8 := uint8(u64)
	return u8
}
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("string to int error", err)
		return int(UNKNOWN)
	}
	return i
}

func StringToInt32(s string) int32 {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Println("string to int error", err)
		return int32(UNKNOWN)
	}
	return int32(i64)
}

func StringToInt64(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("string to int error", err)
		return int64(UNKNOWN)
	}
	return i64
}

func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func Float64ToString(i float64) string {
	return fmt.Sprintf("%f", i)
}

func StringToFloat32(s string) float64 {
	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Println("string to float32 error", err)
	}
	return f32
}
func StringToFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("string to float64 error", err)
	}
	return f64
}

func Float64ToUint8(s float64) uint8 {
	u8 := uint8(s)
	return u8
}

func Float64ToInt64(s float64) int64 {
	it := int64(s)
	return it
}

func HidePhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + strings.Repeat(SYMBOL_ASTERISK, 4) + phone[len(phone)-4:]
}

func IsHidePhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}
	subMobileId := SubString(phone, 3, 4)
	return subMobileId == strings.Repeat(SYMBOL_ASTERISK, 4)
}

// https://www.cnblogs.com/heris/p/16025741.html
// 返回字符串的长度
func StrLen(str string) int {
	return utf8.RuneCountInString(str)
}
