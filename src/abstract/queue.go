/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

// 基于FIFO策略的集合类型
// 元素处理顺序为其被添加到队列中的顺序
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
	Peek() interface{}
	Iterate() Iterator
}
