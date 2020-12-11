/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/6 19:38
 */

package abstract

type SymbolTable interface {
	Put(key, val interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	Contains(key interface{}) bool
	IsEmpty() bool
	Size() int
	Keys() Iterator
}

type SortedSymbolTable interface {
	Put(key, val interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	Contains(key interface{}) bool
	IsEmpty() bool
	Size() int
	Min() interface{}                    // 最小的键
	Max() interface{}                    // 最大的键
	Floor(key interface{}) interface{}   // 小于等于key的最大键
	Ceiling(key interface{}) interface{} // 大于等于key的最小键
	Rank(key interface{}) int            // 小于key的键的数量
	Select(k int) interface{}            // 排名为k的键
	DeleteMin()
	DeleteMax()
	RangeSize(lo, hi interface{}) int      //[lo..hi]之间键的数量
	RangeKeys(lo, hi interface{}) Iterator //[lo..hi]之间的所有键，已排序
	Keys() Iterator
}
