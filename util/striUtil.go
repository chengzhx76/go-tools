package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"io"
	"log"
	"strconv"
	"strings"
)

//https://www.cnblogs.com/lpgit/p/10632756.html

func IsBlank(str string) bool {
	str = TrimSpace(str)
	return len(str) == 0 || str == SYMBOL_EMPTY
}

func IsHasBlank(strs ...string) bool {
	for _, v := range strs {
		if IsBlank(v) {
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

func NilToBlank(data interface{}) string {
	if IsNil(data) {
		return ""
	}
	val, ok := data.(string)
	if ok {
		return val
	} else {
		log.Println("data not string type")
	}
	return ""
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
		args += key + "=" + value + "&"
	}
	return TrimSuffix(args, "&")
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

// 从第一位开始截取，返回截取的字符串
// length 截取的个数
func SubBeforeString(s string, length int) string {
	endIndex := length
	if endIndex > len(s) {
		endIndex = len(s)
	}
	return s[:endIndex]
}

// 从最后一位开始截取，返回截取的字符串
// length 截取的个数
func SubAfterString(s string, length int) string {

	endIndex := len(s)
	startIndex := endIndex - length
	if startIndex < 0 {
		startIndex = 0
	}

	return s[startIndex:endIndex]
}

// 从 start 位开始截取，截取 length 个，返回截取的字符串
// length 截取的个数
// ex: util.SubString("0123456789", 4, 2) = 45
func SubString(s string, start, length int) string {

	startIndex := start
	if start < 0 {
		startIndex = 0
	}
	endIndex := len(s)
	if length < 0 {
		endIndex = startIndex
	} else if length <= endIndex {
		endIndex = startIndex + length
	}
	return s[startIndex:endIndex]
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func ValStr(body map[string]interface{}, key string) string {
	valObj := body[key]
	if IsNil(valObj) {
		return SYMBOL_EMPTY
	}
	val, ok := valObj.(string)
	if ok {
		return val
	} else {
		log.Println("<%s> not string type", key)
	}
	return val
}

func ValString(body map[string]string, key string) string {
	valObj := body[key]
	if IsNil(valObj) {
		return ""
	}
	return valObj
}

func ValSlice(body map[string]interface{}, key string) []interface{} {
	valObj := body[key]
	if IsNil(valObj) {
		return nil
	}
	val, ok := valObj.([]interface{})
	if ok {
		return val
	} else {
		log.Println("<%s> not slice type return default val", key)
	}
	return nil
}

func ValFloat64(body map[string]interface{}, key string) float64 {
	valObj := body[key]
	val, ok := valObj.(float64)
	if ok {
		return val
	} else {
		log.Println("<%s> not float64 type return default val 0", key)
	}
	return 0
}

func ValFloat64ToInt32(body map[string]interface{}, key string) int32 {
	valObj := body[key]
	val, ok := valObj.(float64)
	if ok {
		return int32(val)
	} else {
		log.Println("<%s> not float64 type return default val 0", key)
	}
	return 0
}
func ValFloat64ToInt64(body map[string]interface{}, key string) int64 {
	valObj := body[key]
	val, ok := valObj.(float64)
	if ok {
		return int64(val)
	} else {
		log.Println("<%s> not float64 type return default val 0", key)
	}
	return 0
}

func ValUnit8(body map[string]interface{}, key string) uint8 {
	valObj := body[key]
	val, ok := valObj.(float64)
	if ok {
		return uint8(val)
	} else {
		log.Println("<%s> not float64.unit8 type return default val 0", key)
	}
	return 0
}

func ValBool(body map[string]interface{}, key string) bool {
	valObj := body[key]
	val, ok := valObj.(bool)
	if ok {
		return val
	} else {
		log.Println("<%s> not bool type return default val false", key)
	}
	return false
}

func ValMap(body map[string]interface{}, key string) map[string]interface{} {
	valObj := body[key]
	val, ok := valObj.(map[string]interface{})
	if ok {
		return val
	} else {
		log.Println("<%s> not map type return default val nil map", key)
	}
	return nil
}

func InterfaceToMap(data interface{}, defVal ...map[string]interface{}) map[string]interface{} {
	val, ok := data.(map[string]interface{})
	if ok {
		return val
	} else {
		if len(defVal) > 0 && defVal[0] != nil {
			return defVal[0]
		}
		log.Println("InterfaceToMap data not map type return default val nil map")
	}
	return val
}

func InterfaceToString(data interface{}) string {
	val := NilToBlank(data)
	return val
}

func InterfaceToInt(data interface{}, defVal ...int) int {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		log.Println("InterfaceToInt data is nil ret default 0")
		return 0
	}
	val, ok := data.(float64)
	if ok {
		return int(val)
	} else {
		log.Println("InterfaceToInt data not num type ret default 0")
		return 0
	}
}
func InterfaceToInt64(data interface{}, defVal ...int64) int64 {
	if data == nil {
		if len(defVal) > 0 {
			return defVal[0]
		}
		log.Println("InterfaceToInt64 data is nil ret default 0")
		return 0
	}

	val := int64(0)

	switch data.(type) {
	case float64:
		val = int64(data.(float64))
	case string:
		val = StringToInt64(data.(string))
	default:
		log.Println("InterfaceToInt64 data not num type ret default 0")
	}
	return val
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Uint8ToString(i uint8) string {
	return strconv.Itoa(int(i))
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
	u64, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		log.Println("string to uint8 error", err)
	}
	u8 := uint8(u64)
	return u8
}
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("string to int error", err)
	}
	return i
}

func StringToInt32(s string) int32 {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Println("string to int error", err)
	}
	return int32(i64)
}

func StringToInt64(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("string to int error", err)
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

func StringToFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 10)
	if err != nil {
		log.Println("string to float64 error", err)
	}
	return f64
}

func Float64ToUint8(s float64) uint8 {
	u8 := uint8(s)
	return u8
}

func HidePhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}
