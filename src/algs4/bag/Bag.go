package main

import (
	"os"
	"fmt"

	. "algs4/util"
)

/**
* Bag
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

type Node struct {
	item interface{}
	next *Node
}

// stack without delete function
type Bag struct {
	first *Node
	n     int
}

func NewBag() *Bag {
	b := &Bag{}
	return b
}

func (b *Bag) isEmpty() bool { return b.n == 0 }

func (b *Bag) size() int { return b.n }

func (b *Bag) add(item interface{}) {
	oldFirst := b.first
	b.first = &Node{item, nil}
	b.first.next = oldFirst
	b.n++
}

func (b *Bag) iterator() []interface{} {
	ret := []interface{}{}
	cur := b.first
	for cur != nil {
		ret = append(ret, cur.item)
		cur = cur.next
	}
	return ret
}

func main() {
	bag := NewBag()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		bag.add(in.ReadString())
	}
	fmt.Println("size of bag = ", bag.size())
	for _, i := range bag.iterator() {
		fmt.Println(i)
	}
}
