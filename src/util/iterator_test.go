package util

import (
	"fmt"
	"testing"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type sliceBag struct {
	a []interface{}
	n int
}

func (b *sliceBag) add(item interface{}) {
	b.a[b.n] = item
	b.n++
}

func (b *sliceBag) Iterate() Iterator {
	ai := NewSliceIterator(b.a, b.n, false)
	return ai
}

type LinkedListBag struct {
	first *Node
	n     int
}

func (b *LinkedListBag) add(item interface{}) {
	oldFirst := b.first
	b.first = &Node{Item: item}
	b.first.Next = oldFirst
	b.n++
}

func (b *LinkedListBag) Iterate() Iterator {
	ai := NewLinkedListIterator(b.first)
	return ai
}

func TestNewSliceIterator(t *testing.T) {
	b := &sliceBag{a: make([]interface{}, 10)}
	input := []int{5, 6, 7, 8}
	for idx := range input {
		b.add(input[idx])
	}
	iter := b.Iterate()
	idx := 0
	for w := iter.Next(); w != nil; w = iter.Next() {
		if w != input[idx] {
			t.Error("not consistent")
		}
		idx++
	}
	// reuse iterator
	iter.Reset()
	idx = 0
	for w := iter.Next(); w != nil; w = iter.Next() {
		if w != input[idx] {
			t.Error("not consistent")
		}
		idx++
	}
}

type tt struct {
	a int
	b int
}

func (t *tt) print() {
	fmt.Println(t.a, t.b)
}

func TestNewLinkedListIterator(t *testing.T) {
	b := &LinkedListBag{}
	input := []int{5, 6, 7, 8}
	for idx := range input {
		b.add(&tt{input[idx], input[idx] + 5})
	}
	iter := b.Iterate()
	for w := iter.Next(); w != nil; w = iter.Next() {
		fmt.Println(w)
	}
	// reuse iterator
	iter.Reset()
	for w := iter.Next(); w != nil; w = iter.Next() {
		fmt.Println(w)
	}

	// test
	iter = b.Iterate()
	for w := iter.Next(); w != nil; w = iter.Next() {
		iter2 := b.Iterate()
		for v := iter2.Next(); v != nil; v = iter2.Next() {
			if v == w {
				fmt.Println(w, v, w == v, &w, &v)
			}
		}
	}
}
