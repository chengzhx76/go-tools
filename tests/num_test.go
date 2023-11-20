package tests

import (
	"fmt"
	"github.com/chengzhx76/go-tools/util"
	"strings"
	"testing"
)

func TestNum(t *testing.T) {
	t.Log(util.NumToHex(100003948, 10))
	//t.Log(util.NumToHex(100003948, 16))
	//t.Log(util.NumToHex(100003948, 36))
	//t.Log(util.NumToHex(100003948, 38))
	//t.Log(util.NumToHex(100003948, 64))

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
	t.Log(util.SubString("0123456789", 0, -2))
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
