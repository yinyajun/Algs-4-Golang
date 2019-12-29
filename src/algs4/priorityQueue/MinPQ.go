package priorityQueue

import (
	. "util"
)

/**
* Min Priority Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */
type MinPQ struct {
	pq         []Key // pa[0] is not used
	n          int
	comparator Comparator
}

func NewMinPQwithCapAndCom(capacity int, compartor Comparator) *MinPQ {
	pq := &MinPQ{}
	pq.pq = make([]Key, capacity+1)
	pq.comparator = compartor
	return pq
}

func NewMinPQwithCap(capacity int) *MinPQ {
	return NewMinPQwithCapAndCom(capacity, nil)
}

func NewMinPQ() *MinPQ {
	return NewMinPQwithCap(1)
}

func NewMinPQwithArray(keys []Key) *MinPQ {
	n := len(keys)

	pq := NewMinPQwithCap(n)
	pq.n = n
	for idx, key := range keys {
		pq.pq[idx+1] = key
	}
	pq.Heapify()
	return pq
}

func (m *MinPQ) IsEmpty() bool { return m.n == 0 }

func (m *MinPQ) Size() int { return m.n }

func (m *MinPQ) Min() Key {
	if m.IsEmpty() {
		panic("Min: MinPQ underflows")
	}
	return m.pq[1]
}

// helper function to double the Size of the heap array
func (m *MinPQ) resize(capacity int) {
	tmp := make([]Key, capacity)
	for i := 1; i <= m.n; i++ {
		tmp[i] = m.pq[i]
	}
	m.pq = tmp
}

func (m *MinPQ) Insert(x Key) {
	// double Size of array if necessary
	if m.n == len(m.pq)-1 {
		m.resize(2 * len(m.pq))
	}

	m.n++
	m.pq[m.n] = x
	m.swim(m.n)
	if !m.isMinHeap() {
		panic("Insert: Insert failed")
	}
}

func (m *MinPQ) DelMin() Key {
	if m.IsEmpty() {
		panic("MinPQ underflows")
	}
	min := m.pq[1]
	m.exch(m.n, 1)
	m.n--
	m.sink(1)
	m.pq[m.n+1] = nil
	if m.n > 0 && m.n == (len(m.pq)-1)/4 {
		m.resize(len(m.pq) / 2)
	}
	if !m.isMinHeap() {
		panic("DelMin: DelMin failed")
	}
	return min
}

func (m *MinPQ) Heapify() {
	for k := m.n / 2; k >= 1; k-- {
		m.sink(k)
	}
}

func (m *MinPQ) HeapAdjust(key Key) {
	m.pq[1] = key
	m.sink(1)
}

/***************************************************************************
 * Helper functions to restore the heap invariant.
 ***************************************************************************/
func (m *MinPQ) swim(k int) {
	// parent index = k/2
	// 保证父节点存在,如果小于父节点，就和父节点交换
	for k > 1 && m.great(k/2, k) {
		m.exch(k, k/2)
		k = k / 2
	}
}

func (m *MinPQ) sink(k int) {
	// left child index = 2k, right child index =2k+1
	// 保证子节点存在，找到最小的子节点，如果大于之，则交换
	for 2*k <= m.n {
		// get Min child between left and right child
		j := 2 * k
		if j < m.n && m.great(j, j+1) {
			j = j + 1
		}
		if !m.great(k, j) {
			break
		}
		m.exch(k, j)
		k = j
	}
}

/***************************************************************************
 * Helper functions for compares and swaps.
 ***************************************************************************/
func (m *MinPQ) great(i, j int) bool {
	if m.comparator == nil {
		return Great(m.pq[i], m.pq[j])
	}
	return m.comparator.Compare(m.pq[i], m.pq[j]) > 0
}

func (m *MinPQ) exch(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
}

// is PQ[1..n] a Min heap?
func (m *MinPQ) isMinHeap() bool {
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
	return m.isMinHeapOrdered(1)
}

// is subtree of PQ[1..n] rooted at k a Min heap?
func (m *MinPQ) isMinHeapOrdered(k int) bool {
	if k > m.n {
		return true
	}
	left := 2 * k
	right := left + 1
	if left <= m.n && m.great(k, left) {
		return false
	}
	if right <= m.n && m.great(k, right) {
		return false
	}
	return m.isMinHeapOrdered(left) && m.isMinHeapOrdered(right)
}

func (m *MinPQ) ExtractItem() interface{} {
	return m.DelMin()
}

func (m *MinPQ) Iterate() Iterator {
	cop := NewMinPQwithCap(len(m.pq) - 1)
	for i := 1; i <= m.n; i++ {
		cop.Insert(m.pq[i])
	}
	return NewCopyIterator(cop)
}
