package main

import (
	"fmt"
	"reflect"
	"algs4/util"
	"os"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

type FixedCapacityStack struct {
	a []interface{}
	N int
}

func NewFixedCapacityStrings(cap int) *FixedCapacityStack {
	return &FixedCapacityStack{
		a: make([]interface{}, cap),
	}
}

func (m *FixedCapacityStack) isEmpty() bool {
	return m.N == 0
}

func (m *FixedCapacityStack) size() int {
	return m.N
}

func (m *FixedCapacityStack) setArray(idx int, item interface{}) {
	slice := reflect.ValueOf(m.a)
	slice.Index(idx).Set(reflect.ValueOf(item))
}

func (m *FixedCapacityStack) push(item interface{}) {
	m.setArray(m.N, item)
	m.N++
}

func (m *FixedCapacityStack) pop() interface{} {
	m.N--
	ret := m.a[m.N]
	return ret
}

func main() {
	s := NewFixedCapacityStrings(100)
	in := util.NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			s.push(item)
		} else if !s.isEmpty() {
			fmt.Print(s.pop(), " ")
		}
	}
	fmt.Println("(", s.size(), "left on stack)")
}
