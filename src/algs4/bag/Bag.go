package bag

import . "util"

/**
* Bag
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

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
	b.first = &Node{Item: item, Next: nil}
	b.first.Next = oldFirst
	b.n++
}

func (b *Bag) Iterate() Iterators { return NewLinkedListIterator(b.first) }
