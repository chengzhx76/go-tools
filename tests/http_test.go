package tests

import (
	"github.com/chengzhx76/go-tools/util"
	"os"
	"testing"
)

func Test_get(t *testing.T) {
	authorization := "WECHATPAY2-SHA256-RSA2048 mchid=\"1625698386\",nonce_str=\"9wxBEcgiwhBbPiFD2xi0sBhxd1yZD79M\",timestamp=\"1683700588\",serial_no=\"2FB52D4E295404CF583FDBD07A78E2F461214308\",signature=\"xquu+Dmf9ON07pP5e3vveJS6rCT6sV5nTPuqbWZ+apaulfvMgwsRSnyDJj/D5aNqcprQTjSU8ps0xmtbq0nlC1UkE27CEef6oOvyj24UwHl1OOPAk+cXFVW1m0dBEwOpMYtuwkbY9dFRhrm/tk9rIshdbfR6SAXgZpG/3u26ZZ12/JssX7mHd8ePzOo0SFEj2J342uSu+kbhDaLVoJITEPVt8rxIizrTd0A2URqB3dmtoD35p2EP94YDYpNDub0IMC7Ok9g4vdnBiepPUPA65P6n29Gd0Xk+sPdbzbYn3kVJqseffczniBo+oVaPBR4/y89nTjjYSecq/r0UOPVi2A==\""
	token := "FAqQGco2srX8115MIMOeFaGsLXQXPjepXVaRYtEKjMxUMttRL3nota1hfxkxDOW"

	url := "https://api.mch.weixin.qq.com/v3/billdownload/file?token=" + token
	header := make(map[string]string, 0)
	header["Accept"] = "*/*"
	header["Content-Type"] = "application/json"
	header["Authorization"] = authorization
	fileRes, err := util.GetRequestByte(url, header)
	if err != nil {
		t.Log(err)
	}

	err = os.WriteFile("output.xlsx", fileRes, 0666)
	if err != nil {
		t.Log(err)
	}

	/*out, err := os.Create("output.xlsx")
	defer out.Close()
	_, err = io.Copy(out, fileRes)*/
}
