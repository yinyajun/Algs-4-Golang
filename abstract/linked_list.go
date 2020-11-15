/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

type Node interface {
	Next() Node
	SetNext(Node)
	Value() interface{}
	SetValue(interface{})
}
