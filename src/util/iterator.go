package util

import "fmt"

type Iterator interface {
	Next() (interface{}, bool)
}

func PrintIterator(it Iterator) {
	Map(it, Print)
}

func Map(it Iterator, f func(interface{})) {
	for val, hasNext := it.Next(); hasNext; val, hasNext = it.Next() {
		f(val)
	}
}

func Print(i interface{}) {
	fmt.Println(i)
}
