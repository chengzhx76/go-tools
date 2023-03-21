package tests

import (
	"fmt"
	"github.com/chengzhx76/go-tools/util"
	"strings"
	"testing"
)

func TestNum(t *testing.T) {

	//t.Log(util.From10To10Str(1655879013569))
	//t.Log(util.From10StrTo10("8344120986430"))
	//t.Log(util.From10To10Str(1655879085094))

	t.Log(util.StringToFloat64("39.956886460191924"))

	//t.Log(util.From10To64(1655879085094))
	//t.Log(util.From64To10("o6a5qMC"))
	t.Log("==============================")
	//t.Log(util.From10To38_2(1655879085094))
	//t.Log(util.From38To10_2("eh_auwd8"))
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
