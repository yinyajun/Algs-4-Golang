package main

import (
	"fmt"
	"time"
)

type iterator interface {
	reset()
	init()
	hasNext() bool
	next() interface{}
}

type arrayIterator struct {
	first    int
	cur      int
	_hasNext func() bool
	_next    func() interface{}
}

func (a *arrayIterator) init() {
	a.first = 0
}

func (a *arrayIterator) reset() {
	a.cur = a.first
}

func (a *arrayIterator) hasNext() bool {
	return a._hasNext()
}

func (a *arrayIterator) next() interface{} {
	return a._next()
}

type bag struct {
	a []interface{}
	n int
}

func (b *bag) add(item interface{}) {
	b.a[b.n] = item
	b.n++
}

func (b *bag) Iterate() iterator {
	ai := &arrayIterator{}
	ai.init()
	ai.reset()
	ai._hasNext = func() bool {
		if ai.cur < b.n {
			return true
		} else {
			return false
		}
	}
	ai._next = func() interface{} {
		if ai.hasNext() {
			ret := b.a[ai.cur]
			ai.cur++
			return ret
		}
		return nil
	}
	return ai
}

func main() {
	b := &bag{a: make([]interface{}, 10)}
	b.add(5)
	b.add(6)
	b.add(7)
	b.add(8)
	//k := b.Iterate()
	//it := k.makeGenerator()

	it := b.Iterate()
	for w := it.next(); w != nil; w = it.next() {
		fmt.Println(w)
		time.Sleep(100 * time.Millisecond)
	}

	//for w := it.next(); w != nil; w = it.next() {
	//	fmt.Println(w)
	//	time.Sleep(500 * time.Millisecond)
	//}
	//it = k.makeGenerator()
	//it = b.iterate()
	it.reset()
	for w := it.next(); w != nil; w = it.next() {
		fmt.Println(w)
		time.Sleep(100 * time.Millisecond)
	}

	it.reset()
	for v := it.next(); v != nil; {
		fmt.Println(v)
		time.Sleep(500 * time.Millisecond)
	}

}
