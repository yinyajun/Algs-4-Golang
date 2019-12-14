package util

import "fmt"

/**
* iterator for collections
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Iterate interface {
	Iterate() Iterators
}

type Iterators interface {
	Reset()
	HasNext() bool
	Next() interface{}
}

func PrintIterators(it Iterators) {
	for i := it.Next(); i != nil; i = it.Next() {
		fmt.Println(i)
	}
}

/**
* Iterator for slice
 */
type SliceIterator struct {
	init     int
	cur      int
	slice    []interface{}
	size     int
	reverse  bool
	_hasNext func() bool
	_next    func() interface{}
}

func NewSliceIterator(slice []interface{}, size int, reverse bool) *SliceIterator {
	s := &SliceIterator{slice: slice, reverse: reverse}
	if reverse {
		s.size = size
		s.init = size - 1
		if s.init < 0 {
			s.init = 0
		}
	} else {
		s.init = 0
	}
	s.Reset()
	s.SetHasNextFunc()
	s.SetNextFunc()
	return s
}
func (s *SliceIterator) Reset()            { s.cur = s.init }
func (s *SliceIterator) HasNext() bool     { return s._hasNext() }
func (s *SliceIterator) Next() interface{} { return s._next() }
func (s *SliceIterator) SetHasNextFunc() {
	s._hasNext = func() bool {
		if s.reverse {
			return s.cur >= 0
		}
		return s.cur < len(s.slice)
	}
}
func (s *SliceIterator) SetNextFunc() {
	s._next = func() interface{} {
		if s.HasNext() {
			ret := s.slice[s.cur]
			if s.reverse {
				s.cur--
			} else {
				s.cur++
			}
			return ret
		}
		return nil
	}
}

/**
* Iterator for linked list
 */
type LinkedListIterator struct {
	init     *Node
	cur      *Node
	_hasNext func() bool
	_next    func() interface{}
}

func NewLinkedListIterator(first *Node) *LinkedListIterator {
	l := &LinkedListIterator{init: first}
	l.Reset()
	l.SetNextFunc()
	l.SetHasNextFunc()
	return l
}
func (l *LinkedListIterator) Reset()            { l.cur = l.init }
func (l *LinkedListIterator) HasNext() bool     { return l._hasNext() }
func (l *LinkedListIterator) Next() interface{} { return l._next() }
func (l *LinkedListIterator) SetHasNextFunc() {
	l._hasNext = func() bool { return l.cur != nil }
}
func (l *LinkedListIterator) SetNextFunc() {
	l._next = func() interface{} {
		if l.HasNext() {
			ret := l.cur.Item
			l.cur = l.cur.Next
			return ret
		}
		return nil
	}
}

/**
* Iterator for collections which could be consumed once, e.g. Max Priority Queue
* Since get data operation has side effects (will change data structure)
 */
type OnceCollect interface {
	IsEmpty() bool
	ExtractItem() interface{}
}

type OnceIterator struct {
	collect  OnceCollect
	_hasNext func() bool
	_next    func() interface{}
}

func NewOnceIterator(c OnceCollect) *OnceIterator {
	o := &OnceIterator{collect: c}
	o.SetHasNextFunc()
	o.SetNextFunc()
	return o
}

func (o *OnceIterator) Reset()            { panic("Reset is not supported in OnceIterator") }
func (o *OnceIterator) HasNext() bool     { return o._hasNext() }
func (o *OnceIterator) Next() interface{} { return o._next() }
func (o *OnceIterator) SetHasNextFunc() {
	o._hasNext = func() bool { return !o.collect.IsEmpty() }
}
func (o *OnceIterator) SetNextFunc() {
	o._next = func() interface{} {
		if o.HasNext() {
			return o.collect.ExtractItem()
		}
		return nil
	}
}
