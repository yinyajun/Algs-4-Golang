package main

import (
	"fmt"
	"os"
	"strconv"

	. "algs4/util"
)

/**
* Max Priority Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

//type Key interface{}

type pq struct {
	pq         []Key // pq[0] is not used
	n          int
	comparator Comparator
}

func NewMaxPQwithCapAndCom(capacity int, compartor Comparator) *pq {
	pq := &pq{}
	pq.pq = make([]Key, capacity+1)
	pq.comparator = compartor
	return pq
}

func NewMaxPQwithCap(capacity int) *pq {
	return NewMaxPQwithCapAndCom(capacity, nil)
}

func NewMaxPQ() *pq {
	return NewMaxPQwithCap(1)
}

func NewMaxPQwithArray(keys []Key) *pq {
	n := len(keys)
	pq := NewMaxPQwithCap(n + 1)
	for idx, key := range keys {
		pq.pq[idx+1] = key
	}
	pq.heapify()
	return pq
}

func (m *pq) isEmpty() bool {
	return m.n == 0
}

func (m *pq) size() int {
	return m.n
}

func (m *pq) max() Key {
	if m.isEmpty() {
		panic("max: pq underflows")
	}
	return m.pq[1]
}

// helper function to double the size of the heap array
func (m *pq) resize(capacity int) {
	tmp := make([]Key, capacity)
	for i := 1; i <= m.n; i++ {
		tmp[i] = m.pq[i]
	}
	m.pq = tmp
}

func (m *pq) insert(x Key) {
	// double size of array if necessary
	if m.n == len(m.pq)-1 {
		m.resize(2 * len(m.pq))
	}

	m.n++
	m.pq[m.n] = x
	m.swim(m.n)
}

func (m *pq) delMax() Key {
	if m.isEmpty() {
		panic("maxPQ underflows")
	}
	max := m.pq[1]
	m.exch(m.n, 1)
	m.n--
	m.sink(1)
	m.pq[m.n+1] = nil
	if m.n > 0 && m.n == (len(m.pq)-1)/4 {
		m.resize(len(m.pq) / 2)
	}
	return max
}

func (m *pq) heapify() {
	for k := m.n / 2; k >= 1; k-- {
		m.sink(k)
	}
}

func (m *pq) heapAdjust(key Key) {
	m.pq[1] = key
	m.sink(1)
}

/***************************************************************************
 * Helper functions to restore the heap invariant.
 ***************************************************************************/
func (m *pq) swim(k int) {
	// parent index = k/2
	// 保证父节点存在,如果大于等于父节点，就和父节点交换
	for k > 1 && m.less(k/2, k) {
		m.exch(k, k/2)
		k = k / 2
	}
}

func (m *pq) sink(k int) {
	// left child index = 2k, right child index =2k+1
	// 保证子节点存在，找到最大的子节点，如果小于之，则交换
	for 2*k <= m.n {
		// get max child between left and right child
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
func (m *pq) less(i, j int) bool {
	if m.comparator == nil {
		return Compare(m.pq[i], m.pq[j])
	}
	return m.comparator.Compare(m.pq[i], m.pq[j]) < 0
}

func (m *pq) exch(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
}

// is pq[1..n] a max heap?
func (m *pq) isMaxHeap() bool {
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

// is subtree of pq[1..n] rooted at k a max heap?
func (m *pq) isMaxHeapOrdered(k int) bool {
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

func main() {
	pq := NewMaxPQ()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			k, _ := strconv.Atoi(item)
			pq.insert(k)
		} else if !pq.isEmpty() {
			fmt.Println(pq.delMax(), " ")
		}
	}
	fmt.Println("(", pq.size(), " left on pq")
}
