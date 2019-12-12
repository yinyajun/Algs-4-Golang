package util

import (
	"fmt"
	"strings"
)

type Generator func() (bool, interface{})

func (g Generator) String() string {
	if g == nil {
		return "<nil>"
	}
	ret := strings.Builder{}
	for hasNext, val := g(); hasNext; hasNext, val = g() {
		ret.WriteString(val.(string))
		ret.WriteString("\n")
	}
	return ret.String()
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
