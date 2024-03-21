package tests

import (
	"github.com/chengzhx76/go-tools/util"
	"testing"
)

func Test_file(t *testing.T) {
	dir, name := util.GetDirAndFileName("/home/cheng/trust_gateway/web/static-ddbook/avatar/chengguangcan.png")
	t.Log(dir)
	t.Log(name)
}

func Test_DecodeLastRuneInString(t *testing.T) {
	name := "张三"
	//firstLetter, _ := utf8.DecodeLastRuneInString(name)
	//firstLetter, _ := utf8.DecodeLastRuneInString(name)
	t.Log(util.SubString(name, 0, -1))
}
