package asset

import (
	"errors"
	"github.com/chengzhx76/go-tools/util"
)

type Asset struct {
	Err error
}

func NewAsset() *Asset {
	return &Asset{}
}

func (as *Asset) HasError() bool {
	return as.Err != nil
}

func (as *Asset) IsBlank(str string, err error) {
	if as.Err == nil {
		as.Err = IsBlank(str, err)
	}
}

func (as *Asset) NotBlank(str string, err error) {
	if as.Err == nil {
		as.Err = NotBlank(str, err)
	}
}

func (as *Asset) IsBlankStr(str string, msg string) {
	if as.Err == nil {
		as.Err = IsBlankStr(str, msg)
	}
}

func (as *Asset) NotBlankStr(str string, msg string) {
	if as.Err == nil {
		as.Err = NotBlankStr(str, msg)
	}
}

func (as *Asset) IsNil(obj interface{}, err error) {
	if as.Err == nil {
		as.Err = IsNil(obj, err)
	}
}
func (as *Asset) NotNil(obj interface{}, err error) {
	if as.Err == nil {
		as.Err = NotNil(obj, err)
	}
}
func (as *Asset) IsNilStr(obj interface{}, msg string) {
	if as.Err == nil {
		as.Err = IsNilStr(obj, msg)
	}
}
func (as *Asset) NotNilStr(obj interface{}, msg string) {
	if as.Err == nil {
		as.Err = NotNilStr(obj, msg)
	}
}
func (as *Asset) IsTrue(bl bool, err error) {
	if as.Err == nil {
		as.Err = IsTrue(bl, err)
	}
}
func (as *Asset) NotTrue(bl bool, err error) {
	if as.Err == nil {
		as.Err = NotTrue(bl, err)
	}
}

func (as *Asset) IsTrueStr(bl bool, msg string) {
	if as.Err == nil {
		as.Err = IsTrueStr(bl, msg)
	}
}
func (as *Asset) NotTrueStr(bl bool, msg string) {
	if as.Err == nil {
		as.Err = NotTrueStr(bl, msg)
	}
}
func (as *Asset) IsEmpty(objs []interface{}, err error) {
	if as.Err == nil {
		as.Err = IsEmpty(objs, err)
	}
}
func (as *Asset) NotIsEmpty(objs []interface{}, err error) {
	if as.Err == nil {
		as.Err = NotIsEmpty(objs, err)
	}
}
func (as *Asset) IsEmptyStr(objs []interface{}, msg string) {
	if as.Err == nil {
		as.Err = IsEmptyStr(objs, msg)
	}
}
func (as *Asset) NotIsEmptyStr(objs []interface{}, msg string) {
	if as.Err == nil {
		as.Err = NotIsEmptyStr(objs, msg)
	}
}

// =======================================================

func IsBlankStr(str string, msg string) error {
	return IsBlank(str, errors.New(msg))
}

func NotBlankStr(str string, msg string) error {
	return NotBlank(str, errors.New(msg))
}

func IsNilStr(obj interface{}, msg string) error {
	return IsNil(obj, errors.New(msg))
}

func NotNilStr(obj interface{}, msg string) error {
	return NotNil(obj, errors.New(msg))
}

func IsTrueStr(bl bool, msg string) error {
	return IsTrue(bl, errors.New(msg))
}

func NotTrueStr(bl bool, msg string) error {
	return NotTrue(bl, errors.New(msg))
}

func IsEmptyStr(objs []interface{}, msg string) error {
	return IsEmpty(objs, errors.New(msg))
}

func NotIsEmptyStr(objs []interface{}, msg string) error {
	return NotIsEmpty(objs, errors.New(msg))
}

// ==================================================

func IsBlank(str string, err error) error {
	if !util.IsBlank(str) {
		return err
	}
	return nil
}

func NotBlank(str string, err error) error {
	if util.IsBlank(str) {
		return err
	}
	return nil
}

func IsNil(str interface{}, err error) error {
	if util.IsNil(str) {
		return err
	}
	return nil
}

func NotNil(obj interface{}, err error) error {
	if util.IsNil(obj) {
		return err
	}
	return nil
}

func IsTrue(bl bool, err error) error {
	if !bl {
		return err
	}
	return nil
}

func NotTrue(bl bool, err error) error {
	if bl {
		return err
	}
	return nil
}

func IsEmpty(objs []interface{}, err error) error {
	if len(objs) != 0 {
		return err
	}
	return nil
}

func NotIsEmpty(objs []interface{}, err error) error {
	if len(objs) == 0 {
		return err
	}
	return nil
}
