/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

type Iterable interface {
	Iterate() Iterator
}

type Iterator interface {
	First()
	HasNext() bool
	Next() interface{}
}

type Ranger interface {
	InitFirstIndex() int
	InitLastIndex() int
	GetIndex(int) int
	GetNextIndex() int
}
