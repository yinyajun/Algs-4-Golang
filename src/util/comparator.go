package util

import (
	"reflect"
)

type Key interface{}

type Comparator interface {
	Compare(i, j interface{}) int
}

func getFiledByName(e interface{}, name string) reflect.Value {
	return reflect.ValueOf(e).Elem().FieldByName(name)
}

func comparator(a, b interface{}, name string) bool {
	if getFiledByName(a, name).IsValid() && getFiledByName(b, name).IsValid() {
		return getFiledByName(a, name).Float() < getFiledByName(b, name).Float()
	}
	panic("comparator: invalid field name.")
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
	case float64:
		return a.(float64) < b.(float64)
	default:
		if reflect.TypeOf(a).Elem().Name() == "Edge" {
			return comparator(a, b, "weight")
		}
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
	case float64:
		return a.(float64) <= b.(float64)
	default:
		if reflect.TypeOf(a).Elem().Name() == "Edge" {
			return comparator(a, b, "weight")
		}
		panic("compare Err: unsupported type")
	}
}

func Great(a, b interface{}) bool {
	return !Leq(a, b)
}
