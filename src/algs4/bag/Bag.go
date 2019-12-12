package bag

import . "util"

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

func (b *Bag) IsEmpty() bool { return b.n == 0 }

func (b *Bag) Size() int { return b.n }

func (b *Bag) Add(item interface{}) {
	oldFirst := b.first
	b.first = &Node{item, nil}
	b.first.next = oldFirst
	b.n++
}

func (b *Bag) Yield() Generator {
	cur := b.first
	return func() (bool, interface{}) {
		if cur != nil {
			ret := cur.item
			cur = cur.next
			return true, ret
		}
		return false, nil
	}
}
