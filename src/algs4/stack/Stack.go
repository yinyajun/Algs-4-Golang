package stack

import . "util"

/**
* stack
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

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
	m.first = &Node{Item: item}
	m.first.Next = oldFirst
	m.N++
}

func (m *Stack) Pop() interface{} {
	if m.IsEmpty() {
		panic("stack underflows")
	}
	// 从栈顶删除元素
	item := m.first.Item
	m.first = m.first.Next
	m.N--
	return item
}

func (m *Stack) Iterate() Iterator { return NewLinkedListIterator(m.first) }
