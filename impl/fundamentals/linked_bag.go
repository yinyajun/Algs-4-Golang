/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package fundamentals

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
)

type LinkedBag struct {
	first *abstract.Node
	n     int
}

func NewLinkedBag() *LinkedBag {
	b := &LinkedBag{}
	return b
}

func (b *LinkedBag) Size() int { return b.n }

func (b *LinkedBag) IsEmpty() bool { return b.n == 0 }

func (b *LinkedBag) Add(item interface{}) {
	newNode := &abstract.Node{Key: item}
	newNode.Next = b.first
	b.first = newNode
	b.n++
}

func (b *LinkedBag) Iterate() abstract.Iterator {
	return utils.NewLinkedListIterator(b.first)
}
