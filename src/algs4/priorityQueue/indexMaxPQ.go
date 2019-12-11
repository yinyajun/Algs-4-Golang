package priorityQueue

import (
	. "util"
)

/**
* Index Max Priority Queue
* 操作索引，而不是实际数据（实际操作pq，而不是keys）
* pq: 二叉堆的位置-> keys索引  qp：keys索引-> 二叉堆的位置
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type IndexPQ struct {
	n    int
	keys []Key
	pq   []int // 1-based indexing
	qp   []int // 1-based re-indexing
}

func NewIndexPQ(maxN int) *IndexPQ {
	if maxN < 0 {
		panic("NewIndexPQ: illegal argument")
	}
	m := &IndexPQ{}
	m.keys = make([]Key, maxN+1)
	m.pq = make([]int, maxN+1)
	m.qp = make([]int, maxN+1)
	for idx := range m.qp {
		m.qp[idx] = -1
	}
	return m
}

func (m *IndexPQ) IsEmpty() bool { return m.n == 0 }

func (m *IndexPQ) Size() int { return m.n }

func (m *IndexPQ) contains(i int) bool { return m.qp[i] != -1 }

// Associate key with index i
func (m *IndexPQ) Insert(i int, key Key) {
	if m.contains(i) {
		panic("Insert: index already in pq ")
	}
	m.n++
	m.keys[i] = key
	m.pq[m.n] = i
	m.qp[i] = m.n
	m.swim(m.n) // 在索引堆中的位置为n，操作索引数组
}

// Returns an index associated with a maximum key.
func (m *IndexPQ) MaxIndex() int {
	if m.n == 0 {
		panic("MaxIndex: priority queue underflows")
	}
	return m.pq[1]
}

// Returns a maximum key.
func (m *IndexPQ) MaxKey() Key {
	if m.n == 0 {
		panic("MaxKey: priority queue underflows")
	}
	return m.keys[m.pq[1]]
}

// Removes a maximum key and returns its associated index.
func (m *IndexPQ) DelMax() int {
	if m.n == 0 {
		panic("DelMax: priority queue underflows")
	}
	max := m.pq[1]
	m.exch(m.n, 1)
	m.n--
	m.sink(1)

	m.pq[m.n+1] = -1
	m.keys[max] = nil
	m.qp[max] = -1
	return max
}

// Returns the key associated with index i
func (m *IndexPQ) KeyOf(i int) Key {
	if !m.contains(i) {
		panic("KeyOf: index not exist")
	}
	return m.keys[i]
}

// Change the key associated with index i to the specified value.
func (m *IndexPQ) ChangeKey(i int, key Key) {
	if !m.contains(i) {
		panic("ChangeKey: index not exist")
	}
	m.keys[i] = key
	m.swim(m.qp[i])
	m.sink(m.qp[i])
}

// Change the key associated with index i to the specified value.
func (m *IndexPQ) Change(i int, key Key) {
	m.ChangeKey(i, key)
}

// Increase the key associated with index i to the specified value.
func (m *IndexPQ) IncreaseKey(i int, key Key) {
	if !m.contains(i) {
		panic("IncreaseKey: index not exist")
	}
	if !Less(m.keys[i], key) {
		panic("IncreaseKey: new key less than original key")
	}
	m.keys[i] = key
	m.swim(m.qp[i])
}

// Decrease the key associated with index i to the specified value.
func (m *IndexPQ) DecreaseKey(i int, key Key) {
	if !m.contains(i) {
		panic("DecreaseKey: index not exist")
	}
	if Leq(m.keys[i], key) {
		panic("DecreaseKey: new key larger than original key")
	}
	m.keys[i] = key
	m.sink(m.qp[i])
}

// Remove the key on the priority queue associated with index {@code i}.
func (m *IndexPQ) Delete(i int) {
	if !m.contains(i) {
		panic("Delete: index not exist")
	}
	idx := m.qp[i]
	m.exch(m.n, idx)
	m.n--
	m.swim(idx)
	m.sink(idx)

	m.keys[i] = nil
	m.qp[i] = -1
	m.pq[m.n+1] = -1

}

/***************************************************************************
 * General helper functions.
 ***************************************************************************/
func (m *IndexPQ) less(i, j int) bool {
	return Less(m.keys[m.pq[i]], m.keys[m.pq[j]])
}

func (m *IndexPQ) exch(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
	m.qp[m.pq[i]], m.qp[m.pq[j]] = i, j
}

/***************************************************************************
* Heap helper functions.
***************************************************************************/
func (m *IndexPQ) swim(k int) {
	// parent index = k/2
	// 保证父节点存在,如果大于等于父节点，就和父节点交换
	for k > 1 && m.less(k/2, k) {
		m.exch(k, k/2)
		k = k / 2
	}
}

func (m *IndexPQ) sink(k int) {
	// left child index = 2k, right child index =2k+1
	// 保证子节点存在，找到最大的子节点，如果小于之，则交换
	for 2*k <= m.n {
		j := 2 * k
		if j < m.n && m.less(j, j+1) {
			j++
		}
		if !m.less(k, j) {
			break
		}
		m.exch(k, j)
		k = j
	}
}

func (m *IndexPQ) Iterator() []int {
	ret := []int{}
	cop := NewIndexPQ(len(m.pq) - 1)
	for i := 1; i <= m.n; i++ {
		cop.Insert(m.pq[i], m.keys[m.pq[i]])
	}
	for !cop.IsEmpty() {
		ret = append(ret, cop.DelMax())
	}
	return ret
}

func (m *IndexPQ) Next() (interface{}, bool) {
	if !m.IsEmpty() {
		return m.DelMax(), !m.IsEmpty()
	}
	return nil, false
}
