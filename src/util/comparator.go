package util

import "reflect"

type Key interface{}

type Comparator interface {
	Compare(i, j interface{}) int
}

func Less(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("compare Err: type mismatch")
	}
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case string:
		return a.(string) < b.(string)
	case float32:
		return a.(float64) < b.(float64)
	default:
		panic("compare Err: unsupported type")
	}
}

func Leq(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("compare Err: type mismatch")
	}
	switch a.(type) {
	case int:
		return a.(int) <= b.(int)
	case string:
		return a.(string) <= b.(string)
	case float32:
		return a.(float64) <= b.(float64)
	default:
		panic("compare Err: unsupported type")
	}
}

func Great(a, b interface{}) bool {
	return !Leq(a, b)
}
