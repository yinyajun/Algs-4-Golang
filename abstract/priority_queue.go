/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/5 22:28
 */

package abstract

type MaxPriorityQueue interface {
	Insert(key interface{})
	Max() interface{}
	DelMax() interface{}
	IsEmpty() bool
	Size() int
}

type MinPriorityQueue interface {
	Insert(key interface{})
	Min() interface{}
	DelMin() interface{}
	IsEmpty() bool
	Size() int
}
