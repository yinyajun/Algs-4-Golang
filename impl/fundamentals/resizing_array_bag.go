/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 16:26
 */

package fundamentals

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
)

const InitCapacity = 8

type ResizingArrayBag struct {
	a []interface{}
	n int // [0, b.n)
}

func NewResizingArrayBag() *ResizingArrayBag {
	return &ResizingArrayBag{make([]interface{}, InitCapacity), 0}
}

func (b *ResizingArrayBag) Size() int { return b.n }

func (b *ResizingArrayBag) IsEmpty() bool { return b.n == 0 }

func (b *ResizingArrayBag) resize(capacity int) {
	utils.Assert(capacity >= b.n, "invalid capacity")
	_copy := make([]interface{}, capacity)
	for i := 0; i < b.n; i++ {
		_copy[i] = b.a[i]
	}
	b.a = _copy
}

func (b *ResizingArrayBag) Add(item interface{}) {
	if b.n == len(b.a) {
		b.resize(2 * b.n)
	}
	b.a[b.n] = item
	b.n++
}

func (b *ResizingArrayBag) Iterate() abstract.Iterator {
	return utils.NewArrayIterator(b.a, 0, b.n, false)
}
