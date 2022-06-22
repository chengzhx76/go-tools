package tests

import (
	"github.com/chengzhx76/go-tools/util"
	"testing"
)

func TestNum(t *testing.T) {

	t.Log(util.From10To10Str(1655879013569))
	t.Log(util.From10StrTo10("8344120986430"))
	t.Log(util.From10To10Str(1655879085094))

	//t.Log(util.From10To64(1655879085094))
	//t.Log(util.From64To10("o6a5qMC"))
	t.Log("==============================")
	//t.Log(util.From10To38_2(1655879085094))
	//t.Log(util.From38To10_2("eh_auwd8"))
}
