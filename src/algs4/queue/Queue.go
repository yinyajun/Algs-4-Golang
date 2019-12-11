package main

import (
	"fmt"
	"os"

	. "util"
)

/**
* Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
	N     int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (m *Queue) isEmpty() bool {
	return m.first == nil
}

func (m *Queue) size() int {
	return m.N
}

func (m *Queue) enqueue(item interface{}) {
	// 向表尾添加元素
	oldLast := m.last
	m.last = &Node{item: item}
	if m.isEmpty() {
		m.first = m.last
	} else {
		oldLast.next = m.last
	}
	m.N++
}

func (m *Queue) dequeue() interface{} {
	if m.isEmpty() {
		panic("queue underflows")
	}
	item := m.first.item
	m.first = m.first.next
	if m.isEmpty() {
		m.last = nil
	}
	m.N--
	return item
}

func main() {
	q := NewQueue()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			q.enqueue(item)
		} else if !q.isEmpty() {
			fmt.Print(q.dequeue(), " ")
		}
	}
	fmt.Println("(", q.size(), "left on queue)")
}
