package util

import (
	"fmt"
	"strconv"
	"strings"
)

type Generator func() (bool, interface{})

func (g Generator) String() string {
	if g == nil {
		return "<nil>"
	}
	ret := strings.Builder{}
	for hasNext, val := g(); hasNext; hasNext, val = g() {
		switch val.(type) {
		case string:
			ret.WriteString(val.(string))
		case int:
			ret.WriteString(strconv.Itoa(val.(int)))
		case float32:
			ret.WriteString(strconv.FormatFloat(float64(val.(float32)), 'E', -1, 32))
		}
		ret.WriteString("\n")
	}
	return ret.String()
}

func (g Generator) DeepCopy() Generator {
	it := []interface{}{}
	for hasNext, val := g(); hasNext; hasNext, val = g() {
		it = append(it, val)
	}
	i := 0
	return func() (bool, interface{}) {
		if i < len(it) {
			ret := it[i]
			i++
			return true, ret
		}
		return false, nil
	}
}

type Iterator interface {
	Yield() Generator
}

// mimic map function in Python
func Map(it Iterator, f func(interface{})) {
	generator := it.Yield()
	for hasNext, val := generator(); hasNext; hasNext, val = generator() {
		f(val)
	}
}

func MapIterator(it Iterator, f func(interface{})) {
	Map(it, f)
}

func PrintIterator(it Iterator) {
	Map(it, func(v interface{}) { fmt.Println(v) })
}
