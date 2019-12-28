package sort

import (
	"reflect"
)

type topK struct {
	lessFunc func(i, j int) bool

	data      interface{}
	dataValue reflect.Value
	size      int
	exch      func(i, j int)
}

func NewTopK(data interface{}, k int, comp func(i, j int) bool) *topK {
	m := &topK{}
	m.lessFunc = comp
	m.size = k
	m.exch = reflect.Swapper(data)
	m.dataValue = reflect.ValueOf(data)

	m.calcTopK()
	return m
}

func (m *topK) GetTopK() interface{} {
	return m.dataValue.Slice(0, m.size).Interface()
}

func (m *topK) calcTopK() {
	s := m.dataValue

	// s[0, size) is heapified and keeps top size elements
	if s.Len() == 0 || m.size <= 0 {
		m.size = 0
	} else if s.Len() <= m.size {
		m.size = s.Len()
		m.heapify(m.size)
	} else { // s.Len() > m.size
		m.heapify(m.size)
		for o := m.size; o < s.Len(); o++ {
			if m.lessFunc(0, o) {
				m.heapAdjust(o)
			}
		}
	}
	// heap sort on heapified slice, smallest last.
	for idx := m.size - 1; idx > 0; idx-- {
		m.exch(idx, 0)
		m.siftDown(0, idx)
	}
}

// heap property on data[lo, hi)
func (m *topK) siftDown(lo, hi int) {
	root := lo
	for 2*root+1 < hi {
		child := 2*root + 1
		if child+1 < hi && m.lessFunc(child+1, child) {
			child ++
		}
		if !m.lessFunc(child, root) {
			break
		}
		m.exch(child, root)
		root = child
	}
}

// Build heap on data[0, size)
func (m *topK) heapify(size int) {
	for i := (m.size - 1 - 1) / 2; i >= 0; i-- {
		m.siftDown(i, size)
	}
}

// m.data[0, size) is heapified
// pop heap[0] and adjust heap
func (m *topK) heapAdjust(k int) {
	if k < m.size {
		panic("invalid k")
	}
	m.exch(k, 0) // heap[0] = data[j]
	m.siftDown(0, m.size)
}

type FeedData struct {
	Relateid int64   `json:"relateid"`
	Score    float64 `json:"score"`
	FeedType int     `json:"type"`

	Model            string  `json:"model,omitempty"`
	Authorid         int64   `json:"authorid,omitempty"`
	Explain          string  `json:"explain,omitempty"`
	RelativeAnchorId string  `json:"relative_anchor_id,omitempty"`
	RelativeScore    float64 `json:"relative_score,omitempty"`
	SimilarScore     string  `json:"similar_score,omitempty"`
}

type TopKSorter struct {
	heap []*FeedData
	size int
}

func NewTopKSorter(size int) *TopKSorter {
	heapRet := []*FeedData{}
	this := &TopKSorter{heap: heapRet, size: size}
	return this
}

func (m *TopKSorter) GetTopK(data []*FeedData) []*FeedData {
	if len(data) == 0 || m.size <= 0 {
		return data
	}
	for _, o := range data {
		if len(m.heap) < m.size { // heap没满
			m.heap = append(m.heap, o)
			if len(m.heap) == m.size {
				m.Heapify(m.size)
			}
		} else if o.Score > m.heap[0].Score {
			m.HeapAdjust(o)
		}
	}
	// 如果heap中没有size个，此时，heap没有堆性质
	ll := m.size
	if len(m.heap) < m.size {
		ll = len(m.heap)
		m.Heapify(len(m.heap))
	}

	// 将heap做heapSort
	for i := ll - 1; i > 0; i-- {
		m.heap[i], m.heap[0] = m.heap[0], m.heap[i]
		m.siftDown(0, i)
	}
	return m.heap
}

func (m *TopKSorter) siftDown(lo, hi int) {
	// 维护heap[lo,hi)的堆性质
	// 小顶堆的siftDown，父节点大于孩子节点，和最小孩子交换
	// index从0开始计数
	// 左孩子：2x+1; 右孩子：2x+2； 父节点:(x-1)/2
	root := lo
	for {
		child := 2*root + 1 // 初始化为左孩子
		if child >= hi {
			break
		}
		// 判断左右孩子，谁更小
		if child+1 < hi && m.heap[child+1].Score < m.heap[child].Score {
			child++
		}
		if m.heap[root].Score <= m.heap[child].Score { // 当lo对应的元素已经小于child的元素时，堆性质完成
			return
		}
		m.heap[root], m.heap[child] = m.heap[child], m.heap[root]
		root = child
	}
}

func (m *TopKSorter) Heapify(size int) {
	// 从第一个非叶节点开始往前siftDown
	for i := (size - 2) / 2; i >= 0; i-- {
		m.siftDown(i, size)
	}
}

func (m *TopKSorter) HeapAdjust(item *FeedData) {
	m.heap[0] = item // if len(heap) ==0 will raise an error
	m.siftDown(0, m.size)
}
