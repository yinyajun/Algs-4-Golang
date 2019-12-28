package priorityQueue

import . "util"

/**
* Index Min Priority Queue
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type IndexMinPQ struct {
	n    int
	keys []Key
	pq   []int
	qp   []int
}

func NewIndexMinPQ(maxN int) *IndexMinPQ {
	if maxN < 0 {
		panic("NewIndexMinPQ: illegal argument")
	}

	m := &IndexMinPQ{}
	m.keys = make([]Key, maxN+1)
	m.pq = make([]int, maxN+1)
	m.qp = make([]int, maxN+1)
	for idx := range m.qp {
		m.qp[idx] = -1
	}
	return m
}

func (m *IndexMinPQ) IsEmpty() bool { return m.n == 0 }

func (m *IndexMinPQ) Size() int { return m.n }

func (m *IndexMinPQ) Contains(i int) bool { return m.qp[i] != -1 }

func (m *IndexMinPQ) Insert(i int, key Key) {
	if m.Contains(i) {
		panic("Insert: index already in pq ")
	}
	m.n++
	m.keys[i] = key
	m.pq[m.n] = i
	m.qp[i] = m.n
	m.swim(m.n)
}

// Returns an index associated with a minimum key.
func (m *IndexMinPQ) MinIndex() int {
	if m.n == 0 {
		panic("MinIdex: priority queue is empty")
	}
	return m.pq[1]
}

func (m *IndexMinPQ) MinKey() Key {
	return m.keys[m.MinIndex()]
}

func (m *IndexMinPQ) DelMin() int {
	if m.n == 0 {
		panic("DelMin: priority queue underflows")
	}
	min := m.pq[1]
	m.exch(1, m.n)
	m.n--
	m.sink(1)

	m.keys[min] = nil
	m.qp[min] = -1
	m.pq[m.n+1] = -1
	return min
}

func (m *IndexMinPQ) KeyOf(i int) Key {
	if !m.Contains(i) {
		panic("KeyOf: index not exist")
	}
	return m.keys[i]
}

func (m *IndexMinPQ) ChangeKey(i int, key Key) {
	if !m.Contains(i) {
		panic("ChangeKey: index not exist")
	}
	m.keys[i] = key
	m.swim(m.qp[i])
	m.sink(m.qp[i])
}

func (m *IndexMinPQ) Change(i int, key Key) {
	m.ChangeKey(i, key)
}

func (m *IndexMinPQ) IncreaseKey(i int, key Key) {
	if !m.Contains(i) {
		panic("IncreaseKey: index not exist")
	}
	if Less(key, m.keys[i]) {
		panic("IncreaseKey: new key less than original key")
	}
	m.keys[i] = key
	m.sink(m.qp[i])
}

func (m *IndexMinPQ) DecreaseKey(i int, key Key) {
	if !m.Contains(i) {
		panic("DecreaseKey: index not exist")
	}
	if Great(key, m.keys[i]) {
		panic("DecreaseKey: new key larger than original key")
	}
	m.keys[i] = key
	m.swim(m.qp[i])
}

func (m *IndexMinPQ) Delete(i int) {
	if !m.Contains(i) {
		panic("Delete: index not exist")
	}
	idx := m.qp[i]
	m.exch(idx, m.n)
	m.n--

	m.swim(idx)
	m.sink(idx)

	m.keys[i] = nil
	m.pq[m.n+1] = -1
	m.qp[i] = -1
}

/***************************************************************************
 * General helper functions.
 ***************************************************************************/
func (m *IndexMinPQ) great(i, j int) bool {
	return Great(m.keys[m.pq[i]], m.keys[m.pq[j]])
}

func (m *IndexMinPQ) exch(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
	m.qp[m.pq[i]], m.qp[m.pq[j]] = i, j
}

/***************************************************************************
* Heap helper functions.
***************************************************************************/
func (m *IndexMinPQ) swim(k int) {
	for k >= 2 && m.great(k/2, k) {
		m.exch(k/2, k)
		k = k / 2
	}
}

func (m *IndexMinPQ) sink(k int) {
	for 2*k <= m.n {
		j := 2 * k
		if j < m.n && m.great(j, j+1) {
			j++
		}
		if !m.great(k, j) {
			break
		}
		m.exch(k, j)
		k = j
	}
}

func (m *IndexMinPQ) ExtractItem() interface{} { return m.DelMin() }

func (m *IndexMinPQ) Iterate() Iterator {
	cop := NewIndexMinPQ(len(m.pq) - 1)
	for i := 1; i <= m.n; i++ {
		cop.Insert(m.pq[i], m.keys[m.pq[i]])
	}
	return NewCopyIterator(cop)
}
