/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/29 20:59
 */

package sorting

import (
	"Algs-4-Golang/abstract"
)

type mergesSort struct {
	*baseSorter
}

func NewMerge() *mergesSort {
	impl := &mergesSort{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *mergesSort) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	aux := make([]int, n)
	s.sort(a, aux, 0, n-1, less)
}

func (s *mergesSort) sort(a, aux []int, lo, hi int, less func(i, j int) bool) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	s.sort(a, aux, lo, mid, less)
	s.sort(a, aux, mid+1, hi, less)
	s.merge(a, aux, lo, mid, hi, less)
}

// merge a[lo:mid], a[mid+1, hi]
func (s *mergesSort) merge(a, aux []int, lo, mid, hi int, less func(i, j int) bool) {
	// copy
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

//func (s *mergesSort) merge2(a, aux []int, lo, mid, hi int, less func(i, j int) bool) {
//	// copy
//	for i := lo; i <= hi; i++ {
//		aux[i] = a[i]
//	}
//
//	i, j := lo, mid+1
//	k := lo
//	for i <= mid && j <= hi {
//		if !less(aux[j], aux[i]) { // aux[i]<=aux[j]
//			a[k] = aux[i]
//			i++
//			k++
//		} else { // aux[i] > aux[j]
//			a[k] = aux[j]
//			j++
//			k++
//		}
//	}
//	for i <= mid {
//		a[k] = aux[i]
//		i++
//		k++
//	}
//}

type advancedMergeSorter struct {
	*baseSorter
	switchSorter abstract.Sorter
}

func NewAdvancedMerge() *advancedMergeSorter {
	impl := &advancedMergeSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	impl.switchSorter = NewInsertion()
	return impl
}

func (s *advancedMergeSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	aux := make([]int, n)
	for i := 0; i < n; i++ {
		aux[i] = a[i]
	}
	s.sort(a, aux, 0, n-1, less)
}

func (s *advancedMergeSorter) sort(a, aux []int, lo, hi int, less func(i, j int) bool) {
	// 优化1：加快小数组的排序速度
	if hi-lo <= 18 {
		s.switchSorter.IndexSort(a[lo:hi+1], less) // attention!
		return
	}

	mid := lo + (hi-lo)/2
	// 优化2：避免复制元素到辅助数组
	s.sort(aux, a, lo, mid, less)
	s.sort(aux, a, mid+1, hi, less)
	// 优化3：检测数组是否有序
	if less(aux[mid+1], aux[mid]) { // aux[mid] > aux[mid + 1]
		s.merge(a, aux, lo, mid, hi, less)
	} else { // aux[mid] <= aux[mid + 1]
		for i := lo; i <= hi; i++ {
			a[i] = aux[i]
		}
	}
}

func (s *advancedMergeSorter) merge(a, aux []int, lo, mid, hi int, less func(i, j int) bool) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
