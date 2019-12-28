package priorityQueue

import (
	. "util"
)

/**
* Max Priority Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type MaxPQ struct {
	pq         []Key // pq[0] is not used
	n          int
	comparator Comparator
}

func NewMaxPQwithCapAndCom(capacity int, compartor Comparator) *MaxPQ {
	pq := &MaxPQ{}
	pq.pq = make([]Key, capacity+1)
	pq.comparator = compartor
	return pq
}

func NewMaxPQwithCap(capacity int) *MaxPQ {
	return NewMaxPQwithCapAndCom(capacity, nil)
}

func NewMaxPQ() *MaxPQ {
	return NewMaxPQwithCap(1)
}

func NewMaxPQwithArray(keys []Key) *MaxPQ {
	n := len(keys)
	pq := NewMaxPQwithCap(n + 1)
	for idx, key := range keys {
		pq.pq[idx+1] = key
	}
	pq.Heapify()
	return pq
}

func (m *MaxPQ) IsEmpty() bool {
	return m.n == 0
}

func (m *MaxPQ) Size() int {
	return m.n
}

func (m *MaxPQ) Max() Key {
	if m.IsEmpty() {
		panic("Max: MaxPQ underflows")
	}
	return m.pq[1]
}

// helper function to double the Size of the heap array
func (m *MaxPQ) resize(capacity int) {
	tmp := make([]Key, capacity)
	for i := 1; i <= m.n; i++ {
		tmp[i] = m.pq[i]
	}
	m.pq = tmp
}

func (m *MaxPQ) Insert(x Key) {
	// double Size of array if necessary
	if m.n == len(m.pq)-1 {
		m.resize(2 * len(m.pq))
	}

	m.n++
	m.pq[m.n] = x
	m.swim(m.n)
	if !m.isMaxHeap() {
		panic("Insert: Insert failed")
	}
}

func (m *MaxPQ) DelMax() Key {
	if m.IsEmpty() {
		panic("MaxPQ underflows")
	}
	max := m.pq[1]
	m.exch(m.n, 1)
	m.n--
	m.sink(1)
	m.pq[m.n+1] = nil
	if m.n > 0 && m.n == (len(m.pq)-1)/4 {
		m.resize(len(m.pq) / 2)
	}
	if !m.isMaxHeap() {
		panic("DelMax: DelMax failed")
	}
	return max
}

func (m *MaxPQ) Heapify() {
	for k := m.n / 2; k >= 1; k-- {
		m.sink(k)
	}
}

func (m *MaxPQ) HeapAdjust(key Key) {
	m.pq[1] = key
	m.sink(1)
}

/***************************************************************************
 * Helper functions to restore the heap invariant.
 ***************************************************************************/
func (m *MaxPQ) swim(k int) {
	// parent index = k/2
	// 保证父节点存在,如果大于父节点，就和父节点交换
	for k > 1 && m.less(k/2, k) {
		m.exch(k, k/2)
		k = k / 2
	}
}

func (m *MaxPQ) sink(k int) {
	// left child index = 2k, right child index =2k+1
	// 保证子节点存在，找到最大的子节点，如果小于之，则交换
	for 2*k <= m.n {
		// get Max child between left and right child
		j := 2 * k
		if j < m.n && m.less(j, j+1) {
			j = j + 1
		}
		if !m.less(k, j) {
			break
		}
		m.exch(k, j)
		k = j
	}
}

/***************************************************************************
 * Helper functions for compares and swaps.
 ***************************************************************************/
func (m *MaxPQ) less(i, j int) bool {
	if m.comparator == nil {
		return Less(m.pq[i], m.pq[j])
	}
	return m.comparator.Compare(m.pq[i], m.pq[j]) < 0
}

func (m *MaxPQ) exch(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
}

// is MaxPQ[1..n] a Max heap?
func (m *MaxPQ) isMaxHeap() bool {
	for i := 1; i <= m.n; i++ {
		if m.pq[i] == nil {
			return false
		}
	}
	for i := m.n + 1; i < len(m.pq); i++ {
		if m.pq[i] != nil {
			return false
		}
	}
	if m.pq[0] != nil {
		return false
	}
	return m.isMaxHeapOrdered(1)
}

// is subtree of MaxPQ[1..n] rooted at k a Max heap?
func (m *MaxPQ) isMaxHeapOrdered(k int) bool {
	if k > m.n {
		return true
	}
	left := 2 * k
	right := left + 1
	if left <= m.n && m.less(k, left) {
		return false
	}
	if right <= m.n && m.less(k, right) {
		return false
	}
	return m.isMaxHeapOrdered(left) && m.isMaxHeapOrdered(right)
}

func (m *MaxPQ) ExtractItem() interface{} {
	return m.DelMax()
}

func (m *MaxPQ) Iterate() Iterator {
	cop := NewMaxPQwithCap(len(m.pq) - 1)
	for i := 1; i <= m.n; i++ {
		cop.Insert(m.pq[i])
	}
	return NewCopyIterator(cop)
}
