/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/5 22:30
 */

package sorting

type heap struct {
	*baseSorter
}

func NewHeap() *heap {
	impl := &heap{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *heap) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)

	s.heapify(a, less)

	for k := n - 1; k > 0; k-- {
		Exch(a, 0, k)
		s.sink(a, 0, k-1, less)
	}
}

func (s *heap) heapify(a []int, less func(i, j int) bool) { //  从最后一个节点的父节点开始
	n := len(a)
	for i := (n - 1 - 1) / 2; i >= 0; i-- {
		s.sink(a, i, n-1, less)
	}
}

// sink on a[lo, hi]
func (s *heap) sink(a []int, lo, hi int, less func(i, j int) bool) {
	root := lo
	for 2*root+1 <= hi {
		j := 2*root + 1
		if j+1 <= hi && less(a[j], a[j+1]) { // find max child
			j++
		}
		if !less(a[root], a[j]) { // a[root]>=a[j]
			break
		}
		Exch(a, root, j)
		root = j
	}
}
