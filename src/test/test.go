package main

import (
	"reflect"
	"fmt"
)

type feeddata struct {
	Score int
	Name  string
}

func doubleSlice(s interface{}) interface{} {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		fmt.Println("The interface is not a slice.")
		return nil
	}

	v := reflect.ValueOf(s)
	newLen := v.Len()
	newCap := (v.Cap() + 1) * 2
	typ := reflect.TypeOf(s).Elem()
	fmt.Println(v, reflect.TypeOf(s).Elem())

	t := reflect.MakeSlice(reflect.SliceOf(typ), newLen, newCap)
	reflect.Copy(t, v)
	return t.Interface()
}

func main() {
	var a []int
	var value reflect.Value = reflect.ValueOf(a)
	fmt.Println(value.Type())
	value = reflect.New(value.Type()).Elem()
	fmt.Println(value.Type())
	value = reflect.AppendSlice(value, reflect.ValueOf([]int{1, 2}))
	value = reflect.AppendSlice(value, reflect.ValueOf([]int{3, 4, 5}))
	fmt.Println(value.Kind(), value.Slice(0, value.Len()).Interface())
}
