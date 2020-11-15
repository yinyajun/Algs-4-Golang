/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

// 背包是一种不支持从中删除元素的集合数据类型
// 目的就是收集元素并迭代遍历所有收集到的元素
// 迭代的顺序不确定
type Bag interface {
	Add(interface{})
	Size() int
	IsEmpty() bool
	Iterate() Iterator
}
