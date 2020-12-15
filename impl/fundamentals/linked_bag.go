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
	first abstract.Node
	n     int
}

func NewLinkedBag() *LinkedBag {
	b := &LinkedBag{}
	return b
}

func (b *LinkedBag) Size() int { return b.n }

func (b *LinkedBag) IsEmpty() bool { return b.n == 0 }

func (b *LinkedBag) Add(item interface{}) {
	newNode := &Node{item, nil, nil}
	newNode.next = b.first
	b.first = newNode
	b.n++
}

func (b *LinkedBag) Iterate() abstract.Iterator {
	return utils.NewLinkedListIterator(b.first)
}

type Node struct {
	key   interface{}
	value interface{}
	next  abstract.Node
}

func NewNode(key, value interface{}, next abstract.Node) *Node {
	return &Node{key, value, next}
}

func (n *Node) Key() interface{} { return n.key }

func (n *Node) SetKey(key interface{}) { n.key = key }

func (n *Node) Value() interface{} { return n.value }

func (n *Node) SetValue(value interface{}) { n.value = value }

func (n *Node) Next() abstract.Node { return n.next }

func (n *Node) SetNext(node abstract.Node) { n.next = node }
