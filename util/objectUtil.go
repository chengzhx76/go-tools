package util

import (
	"reflect"
)


func IsNils(objs ...interface{}) bool {
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

func IsNil(obj interface{}) bool {
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
