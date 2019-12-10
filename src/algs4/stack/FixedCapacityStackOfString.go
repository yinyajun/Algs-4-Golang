package main

import (
	"fmt"
	"os"

	. "algs4/util"
)

/**
* FixedCapStrStack
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type FixedCapStrStack struct {
	a []string
	N int
}

func NewFixedCapStrStack(cap int) *FixedCapStrStack {
	return &FixedCapStrStack{
		a: make([]string, cap),
	}
}

func (m *FixedCapStrStack) isEmpty() bool {
	return m.N == 0
}

func (m *FixedCapStrStack) size() int {
	return m.N
}

func (m *FixedCapStrStack) push(item string) {
	m.a[m.N] = item
	m.N++
}

func (m *FixedCapStrStack) pop() string {
	m.N--
	ret := m.a[m.N]
	m.a[m.N] = ""
	return ret
}

func main() {
	s := NewFixedCapStrStack(100)
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			s.push(item)
		} else if !s.isEmpty() {
			fmt.Print(s.pop() + " ")
		}
	}
	fmt.Println("(", s.size(), "left on stack)")
}
