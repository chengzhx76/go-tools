package tests

import (
	"fmt"
	. "github.com/chengzhx76/go-tools/util"
	"strings"
	"testing"
)

func TestNum(t *testing.T) {
	t.Log(NumToHex(100003948, 10))
	//t.Log(NumToHex(100003948, 16))
	//t.Log(NumToHex(100003948, 36))
	//t.Log(NumToHex(100003948, 38))
	//t.Log(NumToHex(100003948, 64))

	//10899996051
}

func TestNum_a(t *testing.T) {

	intSlice := []int{1, 2, 3, 4, 5}

	first := intSlice[1:]
	fmt.Printf("First element: %dn", first)

}

var (
	uid           = "123"
	platformID    = "156"
	onlineConnNum = 126
)

func Test_b(t *testing.T) {
	ok, ft := formatOut("by.cheng.log WS add user<%s> conn, platform<%d> total online conn num<%d>", uid, platformID, onlineConnNum)
	fmt.Println(ok, ft)
}

func formatOut(args ...interface{}) (bool, string) {
	if len(args) > 1 {
		arg := args[0]
		format, ok := arg.(string)
		if ok {
			return strings.Contains(format, "%"), format
		}
	}
	return false, ""
}

func Test_SubString(t *testing.T) {
	t.Log(SubString("0123456789", 0, -2))
	//s:="0123456789"
	//t.Log(string(s[1:]))
}

func Test_Bu(t *testing.T) {
	a := 663
	fmt.Println(a)
	//前置补0
	fmt.Printf("%03d", a) //9位，不足前面凑0补齐
	fmt.Println("")
	fmt.Printf("%0*d", 9, a) //同上
}

func Test_RoundFloat(t *testing.T) {
	number := 12.3456789

	fmt.Println(RoundFloat(number, 1))
	fmt.Println(RoundFloat(number, 2))
	fmt.Println(RoundFloat(number, 3))
	fmt.Println(RoundFloat(number, 4))
	fmt.Println(RoundFloat(number, 5))

	number = -12.3456789
	fmt.Println(RoundFloat(number, 0))
	fmt.Println(RoundFloat(number, 1))
	fmt.Println(RoundFloat(number, 10))
}

func Test_RoundUpMultiple(t *testing.T) {
	a := int64(663)
	r := RoundUpMultiple(a, 5)
	fmt.Println(r)
}

func Test_RoundUp(t *testing.T) {
	r := RoundUp(-1.11)
	fmt.Println(r)
}
func Test_RoundDown(t *testing.T) {
	r := RoundDown(-1.11)
	fmt.Println(r)
}

func Test_s(t *testing.T) {
	str := "0123456789"
	s := []rune(str)
	fmt.Println(len(s))
	fmt.Println(str[0:len(s)])
}

func Test_NumberFormatCn(t *testing.T) {
	//r := NumberFormatCn(20001) // 2万
	//r := NumberFormatCn(20100) // 2万1千
	//r := NumberFormatCn(20199) // 2万1千
	r := NumberFormatCn(21900) // 2万2千
	//r := SubString2("0123456789", 4, 2) //  = 45 // 从第4索引个开始截取2个
	//r := SubString2("0123456789", 2, -1) //  = 1
	//r := SubString2("0123456789", -1, -1) // = 8
	//r := SubString2("0123456789", -3, 3) // = 789
	//r := SubString2("0123456789", -1, 3) // = 9
	//r := SubString2("0123456789", -3, -3) // = 456
	//r := SubString2("0123456789", 0, -3) // =
	//r := SubString2("0123456789", -3, 0) // = 789
	//r := SubString2("0123456789", 0, 0) // = 0123456789
	fmt.Println(r)
}

func SubString2(str string, start, length int) string {
	s := []rune(str)
	totalLen := len(s)
	if totalLen == 0 {
		return ""
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
		} else if length < 0 {
			return ""
		} else if length == 0 {
			endIndex = totalLen
		}
	}
	if startIndex > totalLen {
		return ""
	}
	if endIndex > totalLen {
		return string(s[startIndex:])
	} else {
		return string(s[startIndex:endIndex])
	}
}
