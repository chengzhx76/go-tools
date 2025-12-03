package util

import (
	"reflect"
)

func IsNils(objs ...any) bool {
	if len(objs) == 1 {
		return IsNil(objs[0])
	} else {
		for _, obj := range objs {
			if IsNil(obj) {
				return true
			}
		}
	}
	return false
}

func IsNil(obj any) bool {
	if obj == nil {
		return true
	}

	switch obj.(type) {
	case string:
		return IsBlank(obj.(string))
	}

	switch reflect.TypeOf(obj).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(obj).IsNil()
	}

	return false
}

// MapSlice 将切片中的每个元素通过映射函数转换为另一种类型
func MapSlice[T any, R any](slice []T, fn func(T) R) []R {
	if slice == nil {
		return nil
	}
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

/*
// 提取所有 Entitlement 的 ID
ids := MapSlice(ents, func(ent *model.Entitlement) string {
    return ent.ID
})

// 其他示例：将整数切片转为字符串切片
nums := []int{1, 2, 3}
strs := MapSlice(nums, func(n int) string {
    return strconv.Itoa(n)
})
// 结果: ["1", "2", "3"]
*/
