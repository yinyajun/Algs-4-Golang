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

// 循环数组
type ResizingArrayQueue struct {
	q     []interface{}
	n     int
	first int //[first, last)
	last  int
}

func NewResizingArrayQueue() *ResizingArrayQueue {
	return &ResizingArrayQueue{q: make([]interface{}, InitCapacity)}
}

func (q *ResizingArrayQueue) Enqueue(item interface{}) {
	// 此时last无法说明元素个数
	if q.n == len(q.q) {
		q.resize(2 * q.n)
	}
	q.q[q.last] = item
	q.n++
	q.last = (q.last + 1) % len(q.q)
}

func (q *ResizingArrayQueue) Dequeue() interface{} {
	utils.Assert(!q.IsEmpty(), "queue underflow")
	item := q.q[q.first]
	q.q[q.first] = nil // I forgot again
	q.n--
	q.first = (q.first + 1) % len(q.q)

	if q.n > 0 && q.n == len(q.q)/4 {
		q.resize(len(q.q) / 2)
	}
	return item
}

func (q *ResizingArrayQueue) resize(capacity int) {
	utils.Assert(capacity > q.n, "invalid capacity")
	_copy := make([]interface{}, capacity)
	for i := 0; i < q.n; i++ {
		_copy[i] = q.q[(i+q.first)%len(q.q)]
	}
	q.q = _copy
	q.first, q.last = 0, q.n
}

func (q *ResizingArrayQueue) IsEmpty() bool { return q.n == 0 }

func (q *ResizingArrayQueue) Size() int { return q.n }

func (q *ResizingArrayQueue) Peek() interface{} {
	utils.Assert(!q.IsEmpty(), "queue underflow")
	return q.q[q.first]
}

func (q *ResizingArrayQueue) Iterate() abstract.Iterator {
	return utils.NewArrayIterator(q.q, q.first, q.n, true)
}
