package main

import (
	"fmt"
	"os"

	. "algs4/util"
)

/**
* stack
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

type Node struct {
	item interface{}
	next *Node
}

type Stack struct {
	first *Node
	N     int
}

func NewStack() *Stack {
	return &Stack{}
}

func (m *Stack) isEmpty() bool {
	return m.first == nil
}

func (m *Stack) size() int {
	return m.N
}

func (m *Stack) push(item interface{}) {
	// 向栈顶添加元素
	oldFirst := m.first
	m.first = &Node{item: item}
	m.first.next = oldFirst
	m.N++
}

func (m *Stack) pop() interface{} {
	if m.isEmpty() {
		panic("stack underflows")
	}
	// 从栈顶删除元素
	item := m.first.item
	m.first = m.first.next
	m.N--
	return item
}

func main() {
	s := NewStack()
	in := NewIn(os.Stdin)
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
