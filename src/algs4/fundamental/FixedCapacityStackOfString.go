package fundamental

/** 
*
* 
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

import (
	"fmt"
	"util/io"
)

type FixedCapacityStackOfStrings struct {
	a []string
	N int
}

func NewFixedCapacityStackOfStrings(cap int) *FixedCapacityStackOfStrings {
	return &FixedCapacityStackOfStrings{
		a: make([]string, cap),
	}
}

func (m *FixedCapacityStackOfStrings) isEmpty() bool {
	return m.N == 0
}

func (m *FixedCapacityStackOfStrings) size() int {
	return m.N
}

func (m *FixedCapacityStackOfStrings) push(item string) {
	m.a[m.N] = item
	m.N++
}

func (m *FixedCapacityStackOfStrings) pop() string {
	m.N--
	ret := m.a[m.N]
	m.a[m.N] = ""
	return ret
}

func EgFCSS() {
	s := NewFixedCapacityStackOfStrings(100)
	in := io.NewIn(io.SplitFunc("words"))
	for !in.IsEmpty() {
		item := in.ReadString()
		if item != "-" {
			s.push(item)
		} else if !s.isEmpty() {
			fmt.Print(s.pop() + " ")
		}
	}
	fmt.Println("(", s.size(), "left on stack)")
}
