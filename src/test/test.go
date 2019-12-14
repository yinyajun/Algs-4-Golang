package main

import (
	"fmt"
	"time"
)

type Iterators interface {
	Reset()
	HasNext() bool
	Next() interface{}
}

type ArrayIterator struct {
	init     int
	cur      int
	array    []interface{}
	_hasNext func() bool
	_next    func() interface{}
}

func NewArrayIterator(array []interface{}) *ArrayIterator {
	return &ArrayIterator{init: 0, array: array}
}

func (a *ArrayIterator) Reset() { a.cur = a.init }

func (a *ArrayIterator) HasNext() bool { return a._hasNext() }

func (a *ArrayIterator) Next() interface{} { return a._next() }

func (a *ArrayIterator) SetHasNextFunc(f func() bool) {
	if f != nil {
		a._hasNext = f
		return
	}
	a._hasNext = func() bool {
		if a.cur < len(a.array) {
			return true
		}
		return false
	}
}

func (a *ArrayIterator) SetNextFunc(f func() interface{}) {
	if f != nil {
		a._next = f
		return
	}
	a._next = func() interface{} {
		if a.HasNext() {
			ret := a.array[a.cur]
			a.cur++
			return ret
		}
		return nil
	}
}

type bag struct {
	a []interface{}
	n int
}

func (b *bag) add(item interface{}) {
	b.a[b.n] = item
	b.n++
}

func (b *bag) Iterate() Iterators {
	ai := NewArrayIterator(b.a)
	ai.Reset()
	ai.SetHasNextFunc(nil)
	ai.SetNextFunc(nil)
	return ai
}

func main() {
	b := &bag{a: make([]interface{}, 10)}
	b.add(5)
	b.add(6)
	b.add(7)
	b.add(8)

	it := b.Iterate()
	for w := it.Next(); w != nil; w = it.Next() {
		fmt.Println(w)
		time.Sleep(100 * time.Millisecond)
	}

	it.Reset()
	for w := it.Next(); w != nil; w = it.Next() {
		fmt.Println(w)
		time.Sleep(100 * time.Millisecond)
	}

}
