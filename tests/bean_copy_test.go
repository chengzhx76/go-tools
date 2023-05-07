package tests

import (
	"encoding/json"
	"github.com/chengzhx76/go-tools/util"
	"testing"
)

func Test_BeanCopy(t *testing.T) {
	sites := []*UserSite{{
		Uid:     "1",
		Site:    "aa",
		Address: "addr",
	}, {
		Uid:     "2",
		Site:    "bb",
		Address: "bddr",
	}}
	user := &User{
		Name:    "张三",
		Age:     18,
		Address: "帝都",
		Attach: &UserAttach{
			Uid:         "1",
			TencentUser: 1,
			Gender:      1,
			FreezeTips:  "find",
		},
		Site: sites,
	}

	userDTO := new(UserDTO)

	err := util.BeanCopy(userDTO, user)

	if err != nil {
		t.Log(err)
	}

	body, err := json.Marshal(userDTO)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(body))
}

type User struct {
	Name    string
	Age     int32
	Address string
	Attach  *UserAttach `json:"attach"`
	Site    []*UserSite `json:"sites"`
}

type UserAttach struct {
	Uid         string `json:"uid"`
	TencentUser uint8  `json:"tu"`
	Gender      uint8  `json:"gender"`
	FreezeTips  string `json:"freezeTips"`
}

type UserSite struct {
	Uid     string `json:"uid"`
	Site    string `json:"site"`
	Address string `json:"address"`
}

type UserDTO struct {
	Name   string         `json:"name"`
	Attach *UserAttachDTO `json:"attach"`
	Site   []*UserSiteDTO `json:"sites"`
}

type UserAttachDTO struct {
	Uid         string `json:"uid"`
	TencentUser uint8  `json:"tu"`
	Gender      uint8  `json:"gender"`
	FreezeTips  string `json:"freezeTips"`
}

type UserSiteDTO struct {
	Uid     string `json:"uid"`
	Site    string `json:"site"`
	Address string `json:"address"`
}
