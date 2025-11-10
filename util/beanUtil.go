package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// 引用 https://github.com/fishyxin/simple-copy-properties
func BeanCopy(dst, src any) (err error) {
	return CopyStructFields(dst, src)
}

// CopyStructFields 将 src 中的导出字段拷贝到 dst，支持可选的字段别名映射。
//   - dst 必须是非 nil 的结构体指针。
//   - src 可以是结构体或结构体指针。
//   - 可选参数 alias（最多传一个 map）：key 为 src 字段名，value 为 dst 字段名。
//     如果未提供 alias，则使用同名字段策略。
func CopyStructFields(dst, src any, alias ...map[string]string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("copy struct fields panic: %v", r)
		}
	}()

	if dst == nil || src == nil {
		return errors.New("dst and src must be non-nil")
	}

	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Pointer || dstVal.IsNil() {
		return errors.New("dst must be a non-nil pointer to a struct")
	}
	dstVal = dstVal.Elem()
	if dstVal.Kind() != reflect.Struct {
		return errors.New("dst must point to a struct")
	}

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Pointer {
		if srcVal.IsNil() {
			return errors.New("src pointer is nil")
		}
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return errors.New("src must be a struct or pointer to struct")
	}

	var aliasMap map[string]string
	if len(alias) > 0 {
		aliasMap = alias[0]
	}

	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcFieldInfo := srcType.Field(i)
		if srcFieldInfo.PkgPath != "" { // 未导出字段跳过
			continue
		}

		srcFieldName := srcFieldInfo.Name
		dstFieldName := srcFieldName
		if aliasMap != nil {
			if mapped, ok := aliasMap[srcFieldName]; ok {
				dstFieldName = mapped
			}
		}

		dstField := dstVal.FieldByName(dstFieldName)
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		srcField := srcVal.Field(i)
		switch {
		case srcField.Type().AssignableTo(dstField.Type()):
			dstField.Set(srcField)
		case srcField.Type().ConvertibleTo(dstField.Type()):
			dstField.Set(srcField.Convert(dstField.Type()))
		default:
			// 不可赋值也不可转换，跳过
		}
	}

	return nil
}

// MapToStruct 将 map[string]any 映射到结构体。
// - src: 源 map，键应与结构体字段的 json 标签或字段名匹配（区分大小写与大小写转换由 json 处理）
// - dst: 目标结构体指针，例如 &UserInfo{}。
// 返回：映射过程中出现的错误。
func MapToStruct(src map[string]any, dst any) error {
	if src == nil {
		return errors.New("MapToStruct: src must not be nil")
	}
	if dst == nil {
		return errors.New("MapToStruct: dst must not be nil")
	}
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("MapToStruct: dst must be a pointer to struct")
	}

	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, dst); err != nil {
		return err
	}
	return nil
}

// MergeStructFields 仅把指定字段从 src 合并到 dst。
//   - dst 必须是非 nil 的结构体指针。
//   - src 可以是结构体或结构体指针。
//   - fields 是需要合并（覆盖）的字段名列表，未列出的字段保持原状。
func MergeStructFields(dst any, src any, fields ...string) error {
	if dst == nil || src == nil {
		return errors.New("dst and src must not be nil")
	}

	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Pointer || dstVal.IsNil() {
		return errors.New("dst must be a non-nil pointer to struct")
	}
	dstVal = dstVal.Elem()
	if dstVal.Kind() != reflect.Struct {
		return errors.New("dst pointer must point to struct")
	}

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Pointer {
		if srcVal.IsNil() {
			return errors.New("src must not be nil pointer")
		}
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return errors.New("src must be struct or struct pointer")
	}

	if len(fields) == 0 {
		return nil // 没有指定字段，直接返回
	}

	fieldSet := make(map[string]struct{}, len(fields))
	for _, name := range fields {
		fieldSet[name] = struct{}{}
	}

	srcType := srcVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		fieldInfo := srcType.Field(i)
		if fieldInfo.PkgPath != "" { // 未导出字段跳过
			continue
		}

		fieldName := fieldInfo.Name
		if _, ok := fieldSet[fieldName]; !ok {
			continue // 不是指定字段，跳过
		}

		srcFieldVal := srcVal.Field(i)
		dstFieldVal := dstVal.FieldByName(fieldName)
		if !dstFieldVal.IsValid() || !dstFieldVal.CanSet() {
			continue
		}
		if !srcFieldVal.Type().AssignableTo(dstFieldVal.Type()) {
			continue
		}

		dstFieldVal.Set(srcFieldVal)
	}

	return nil
}

// example:
// src := UserDTO{
// 	FullName: "Alice Wang",
// 	Mail:     "alice@example.com",
// 	Age:      25,
// }
// dst := &User{}

// // 1. 不传 alias，默认同名字段拷贝（Age 会被拷贝，Name/Email 不会）
// if err := structmerge.CopyStructFields(dst, src); err != nil {
// 	panic(err)
// }

// // 2. 传入 alias，让 FullName -> Name，Mail -> Email
// alias := map[string]string{
// 	"FullName": "Name",
// 	"Mail":     "Email",
// }
// if err := structmerge.CopyStructFields(dst, src, alias); err != nil {
// 	panic(err)
// }
