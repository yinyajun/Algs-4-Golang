/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

type Node interface {
	Key() interface{}
	Value() interface{}
	Next() Node

	SetNext(Node)
	SetValue(interface{})
	SetKey(interface{})
}

type Node2 interface {
	Key() interface{}
	Value() interface{}
	Left() Node2
	Right() Node2
	Size() int

	SetValue(interface{})
	SetKey(interface{})
	SetLeft(Node2)
	SetRight(Node2)
	SetSize(int)
}
