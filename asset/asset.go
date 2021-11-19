package asset

import (
	"errors"
	"go-wechat-pinche/util"
)

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

func IsBlankStr(str string, msg string) error {
	if !util.IsBlank(str) {
		return errors.New(msg)
	}
	return nil
}

func NotBlankStr(str string, msg string) error {
	if util.IsBlank(str) {
		return errors.New(msg)
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

func IsNilStr(str interface{}, msg string) error {
	if !util.IsNil(str) {
		return errors.New(msg)
	}
	return nil
}

func NotNilStr(obj interface{}, msg string) error {
	if util.IsNil(obj) {
		return errors.New(msg)
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

func IsTrueStr(bl bool, msg string) error {
	if !bl {
		return errors.New(msg)
	}
	return nil
}

func NotTrueStr(bl bool, msg string) error {
	if bl {
		return errors.New(msg)
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

func IsEmptyStr(objs []interface{}, msg string) error {
	if len(objs) != 0 {
		return errors.New(msg)
	}
	return nil
}

func NotIsEmptyStr(objs []interface{}, msg string) error {
	if len(objs) == 0 {
		return errors.New(msg)
	}
	return nil
}
