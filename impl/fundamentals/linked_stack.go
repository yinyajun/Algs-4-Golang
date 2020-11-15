/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 19:28
 */

package fundamentals

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
)

type LinkedStack struct {
	first abstract.Node
	n     int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

func (s *LinkedStack) Push(item interface{}) {
	newNode := &Node{item, nil}
	newNode.next = s.first
	s.first = newNode
	s.n++
}

func (s *LinkedStack) Pop() interface{} {
	utils.Assert(!s.IsEmpty(), "stack underflow")
	item := s.first.Value()
	s.first = s.first.Next()
	s.n--
	return item
}

func (s *LinkedStack) Size() int { return s.n }

func (s *LinkedStack) IsEmpty() bool { return s.first == nil }

func (s *LinkedStack) Peek() interface{} {
	utils.Assert(!s.IsEmpty(), "stack underflow")
	return s.first.Value()
}

func (s *LinkedStack) Iterate() abstract.Iterator {
	return utils.NewLinkedListIterator(s.first)
}
