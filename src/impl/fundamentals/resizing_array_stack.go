/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 18:44
 */

package fundamentals

import (
	"abstract"
	"utils"
)

type ResizingArrayStack struct {
	a []interface{}
	n int //[0, b.n)
}

func NewResizingArrayStack() *ResizingArrayStack {
	return &ResizingArrayStack{make([]interface{}, InitCapacity), 0}
}

func (s *ResizingArrayStack) Push(item interface{}) {
	if len(s.a) == s.n {
		s.resize(2 * s.n)
	}
	s.a[s.n] = item
	s.n++
}

func (s *ResizingArrayStack) Pop() interface{} {
	utils.Assert(!s.IsEmpty(), "stack underflow")
	item := s.a[s.n-1]
	s.a[s.n-1] = nil
	s.n--
	// shrink size of array if necessary
	if s.n > 0 && s.n == len(s.a)/4 {
		s.resize(len(s.a) / 2)
	}
	return item
}

func (s *ResizingArrayStack) Peek() interface{} {
	utils.Assert(!s.IsEmpty(), "stack underflow")
	return s.a[s.n-1]
}

func (s *ResizingArrayStack) resize(capacity int) {
	utils.Assert(capacity >= s.n, "invalid capacity")
	_copy := make([]interface{}, capacity)
	for i := 0; i < s.n; i++ {
		_copy[i] = s.a[i]
	}
	s.a = _copy
}

func (s *ResizingArrayStack) IsEmpty() bool { return s.n == 0 }

func (s *ResizingArrayStack) Size() int { return s.n }

func (s *ResizingArrayStack) Iterate() abstract.Iterator {
	return utils.NewArrayIterator(s.a, 0, s.n, false)
}
