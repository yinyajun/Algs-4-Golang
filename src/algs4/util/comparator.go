package util

import "reflect"

type Key interface{}

type Comparator interface {
	Compare(i, j interface{}) int
}

func Compare(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("compare Err: type mismatch")
	}
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case string:
		return a.(string) < b.(string)
	case float32:
		return a.(float32) < b.(float32)
	default:
		panic("compare Err: unsupported type")
	}
}
