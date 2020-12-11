/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/10 21:39
 */

package searching

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

type sequentialSearchST struct {
	first abstract.Node // link list of key-value pair
	n     int
}

type Pair struct {
	Key   interface{}
	Value interface{}
}

func NewSequentialSearchST() *sequentialSearchST { return &sequentialSearchST{} }

func (t *sequentialSearchST) Put(key, val interface{}) {
	utils.AssertF(key != nil, "Key is nil")

	if val == nil {
		t.Delete(key)
		return
	}

	// key already exists
	for x := t.first; x != nil; x = x.Next() {
		if x.Value().(*Pair).Key == key {
			x.SetValue(&Pair{key, val})
			return
		}
	}
	// a new key
	t.first = fundamentals.NewNode(&Pair{key, val}, t.first)
	t.n++
}

func (t *sequentialSearchST) Get(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	for x := t.first; x != nil; x = x.Next() {
		if x.Value().(*Pair).Key == key {
			return x.Value().(*Pair).Value
		}
	}
	return nil
}

func (t *sequentialSearchST) Delete(key interface{}) {
	utils.AssertF(key != nil, "Key is nil")
	t.first = t.delete(t.first, key)
}

// 删除以x开头的链表中的key
func (t *sequentialSearchST) delete(x abstract.Node, key interface{}) abstract.Node {
	if x == nil {
		return nil
	}
	if x.Value().(*Pair).Key == key {
		t.n--
		return x.Next()
	}
	x.SetNext(t.delete(x.Next(), key))
	return x
}

func (t *sequentialSearchST) Contains(key interface{}) bool {
	utils.AssertF(key != nil, "Key is nil")

	return t.Get(key) != nil
}

func (t *sequentialSearchST) IsEmpty() bool { return t.n == 0 }

func (t *sequentialSearchST) Size() int { return t.n }

func (t *sequentialSearchST) Keys() abstract.Iterator {
	queue := fundamentals.NewLinkedQueue()
	for x := t.first; x != nil; x = x.Next() {
		queue.Enqueue(x.Value().(*Pair).Key)
	}
	return queue.Iterate()
}
