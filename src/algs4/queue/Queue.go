package queue

import . "util"

/**
* Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Queue struct {
	first *Node
	last  *Node
	N     int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (m *Queue) IsEmpty() bool {
	return m.first == nil
}

func (m *Queue) Size() int {
	return m.N
}

func (m *Queue) Enqueue(item interface{}) {
	// 向表尾添加元素
	oldLast := m.last
	m.last = &Node{Item: item}
	if m.IsEmpty() {
		m.first = m.last
	} else {
		oldLast.Next = m.last
	}
	m.N++
}

func (m *Queue) Dequeue() interface{} {
	if m.IsEmpty() {
		panic("queue underflows")
	}
	item := m.first.Item
	m.first = m.first.Next
	if m.IsEmpty() {
		m.last = nil
	}
	m.N--
	return item
}

func (m *Queue) Iterate() Iterator {
	return NewLinkedListIterator(m.first)
}
