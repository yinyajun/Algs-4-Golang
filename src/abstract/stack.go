/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

// 基于LIFO策略的集合类型
// 元素处理顺序和其被压入的顺序恰好相反
type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
	Iterate() Iterator
}
