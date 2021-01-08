/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/6 23:10
 */

package searching

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

const InitCapacity = 8

type binarySearchST struct {
	keys   []interface{}
	values []interface{}
	n      int
}

func NewBinarySearchST() *binarySearchST {
	st := &binarySearchST{
		keys:   make([]interface{}, InitCapacity),
		values: make([]interface{}, InitCapacity),
	}
	return st
}

func (t *binarySearchST) resize(capacity int) {
	utils.Assert(capacity >= t.n, "invalid capacity")
	_copyK := make([]interface{}, capacity)
	_copyV := make([]interface{}, capacity)
	for i := 0; i < t.n; i++ {
		_copyK[i] = t.keys[i]
		_copyV[i] = t.values[i]
	}
	t.keys = _copyK
	t.values = _copyV
}

func (t *binarySearchST) Put(key, val interface{}) {
	utils.AssertF(key != nil, "Key is nil")

	if val == nil {
		t.Delete(key)
		return
	}

	i := t.Rank(key)

	// key already in table
	if i < t.n && t.keys[i] == key {
		t.values[i] = val
		return
	}

	// insert new key-value pair
	if t.n == len(t.keys) {
		t.resize(2 * t.n)
	}

	for j := t.n; j > i; j-- {
		t.keys[j] = t.keys[j-1]
		t.values[j] = t.values[j-1]
	}
	t.keys[i] = key
	t.values[i] = val
	t.n++
}

func (t *binarySearchST) Get(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")

	if t.IsEmpty() {
		return nil
	}
	i := t.Rank(key)
	if i < t.n && t.keys[i] == key {
		return t.values[i]
	}
	return nil
}

func (t *binarySearchST) Delete(key interface{}) {
	utils.AssertF(key != nil, "Key is nil")

	i := t.Rank(key)

	if i == t.n || t.keys[i] != key {
		return
	}

	for j := i; j < t.n-1; j++ {
		t.keys[j] = t.keys[j+1]
		t.values[j] = t.values[j+1]
	}
	t.n--
	t.keys[t.n] = nil
	t.values[t.n] = nil

	if t.n > 0 && t.n == len(t.keys)/4 {
		t.resize(len(t.keys) / 2)
	}
}

func (t *binarySearchST) Contains(key interface{}) bool {
	utils.AssertF(key != nil, "Key is nil")

	return t.Get(key) != nil
}

func (t *binarySearchST) IsEmpty() bool { return t.n == 0 }

func (t *binarySearchST) Size() int { return t.n }

// ----------------------------
// ordered symbol table method
// ----------------------------

func (t *binarySearchST) Min() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Min() with empty symbol table")
	return t.keys[0]
}
func (t *binarySearchST) Max() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Max() with empty symbol table")
	return t.keys[t.n-1]
}

// Returns the largest key in this symbol table less than or equal to key
func (t *binarySearchST) Floor(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")

	i := t.Rank(key)
	if i < t.n && t.keys[i] == key { // key in table
		return t.keys[i]
	}
	if i == 0 { // any key in table greater than key
		return nil
	}
	return t.keys[i-1]
}

// Returns the smallest key in this symbol table greater than or equal to key
func (t *binarySearchST) Ceiling(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	i := t.Rank(key)
	if i == t.n { // any key in table less than key
		return nil
	}
	return t.keys[i]
}

// 迭代版本
// 从[lo, hi]区间中搜索比key小的键的数目，值域为[0, n]
func (t *binarySearchST) Rank(key interface{}) int {
	utils.AssertF(key != nil, "Key is nil")

	lo, hi := 0, t.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := utils.CompareTo(t.keys[mid], key)
		if cmp == 0 {
			return mid
		} else if cmp > 0 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
func (t *binarySearchST) Select(k int) interface{} {
	utils.AssertF(k >= 0 && k < t.Size(), "invalid k")
	return t.keys[k]
}

func (t *binarySearchST) DeleteMin() { t.Delete(t.Min()) }

func (t *binarySearchST) DeleteMax() { t.Delete(t.Max()) }

// Returns the number of keys in this symbol table in the specified range.
func (t *binarySearchST) RangeSize(lo, hi interface{}) int {
	utils.AssertF(lo != nil && hi != nil, "invalid lo or hi")

	if utils.Less(hi, lo) {
		return 0
	}
	left := t.Rank(lo)
	right := t.Rank(hi)
	if t.Contains(hi) {
		right += 1
	}
	return right - left
}

// Returns all keys in this symbol table in the given range
func (t *binarySearchST) RangeKeys(lo, hi interface{}) abstract.Iterator {
	utils.AssertF(lo != nil && hi != nil, "invalid lo or hi")

	queue := fundamentals.NewLinkedQueue()

	if utils.Less(hi, lo) {
		return queue.Iterate()
	}

	left := t.Rank(lo)
	right := t.Rank(hi)
	for i := left; i < right; i++ {
		queue.Enqueue(t.keys[i])
	}
	if t.Contains(hi) {
		queue.Enqueue(t.keys[right])
	}
	return queue.Iterate()
}

func (t *binarySearchST) Keys() abstract.Iterator {
	return t.RangeKeys(t.Min(), t.Max())
}
