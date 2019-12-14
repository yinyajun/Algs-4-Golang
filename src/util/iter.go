package util

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Iterators interface {
	Reset()
	HasNext() bool
	Next() interface{}
}

type ArrayIterator struct {
	init     int
	cur      int
	_hasNext func() bool
	_next    func() interface{}
}

func NewArrayIterator() *ArrayIterator {
	return &ArrayIterator{init: 0}
}

func (a *ArrayIterator) Reset() { a.cur = a.init }

func (a *ArrayIterator) HasNext() bool { return a._hasNext() }

func (a *ArrayIterator) Next() interface{} { return a._next() }

func (a *ArrayIterator) SetHasNextFunc(f func() bool) { a._hasNext = f }

func (a *ArrayIterator) SetNextFunc(f func() interface{}) { a._next = f }

type bag struct {
	a []interface{}
	n int
}

func (b *bag) add(item interface{}) {
	b.a[b.n] = item
	b.n++
}

func (b *bag) Iterate() Iterators {
	ai := NewArrayIterator()
	ai.Reset()
	ai.SetHasNextFunc(func() bool {
		if ai.cur < b.n {
			return true
		} else {
			return false
		}
	})
	ai.SetNextFunc(func() interface{} {
		if ai.HasNext() {
			ret := b.a[ai.cur]
			ai.cur++
			return ret
		}
		return nil
	})
	return ai
}
