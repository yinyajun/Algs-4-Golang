package util

import "fmt"

type Generator func() (bool, interface{})

type Iterator interface {
	Yield() Generator
}

// mimic map function in Python
func Map(it Iterator, f func(interface{})) {
	iterator := it.Yield()
	for hasNext, val := iterator(); hasNext; hasNext, val = iterator() {
		f(val)
	}
}

func MapIterator(it Iterator, f func(interface{})) {
	Map(it, f)
}

func PrintIterator(it Iterator) {
	Map(it, func(v interface{}) { fmt.Println(v) })
}
