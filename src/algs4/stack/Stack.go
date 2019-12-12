package stack

import "util"

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

func (m *Stack) IsEmpty() bool {
	return m.first == nil
}

func (m *Stack) Size() int {
	return m.N
}

func (m *Stack) Push(item interface{}) {
	// 向栈顶添加元素
	oldFirst := m.first
	m.first = &Node{item: item}
	m.first.next = oldFirst
	m.N++
}

func (m *Stack) Pop() interface{} {
	if m.IsEmpty() {
		panic("stack underflows")
	}
	// 从栈顶删除元素
	item := m.first.item
	m.first = m.first.next
	m.N--
	return item
}

func (m *Stack) Yield() util.Generator {
	return func() (bool, interface{}) {
		if !m.IsEmpty() {
			return true, m.Pop()
		}
		return false, nil
	}
}
