/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/30 10:04
 */

package sorting

import (
	"Algs-4-Golang/abstract"
	"math/rand"
)

// ----------------------------
// quick sort
// ----------------------------

type quickSorter struct {
	*baseSorter
}

func NewQuick() *quickSorter {
	impl := &quickSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *quickSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	s.sort(a, 0, n-1, less)
}

func (s *quickSorter) sort(a []int, lo, hi int, less func(i, j int) bool) {
	if hi <= lo {
		return
	}
	j := s.partition2(a, lo, hi, less)
	s.sort(a, lo, j-1, less)
	s.sort(a, j+1, hi, less)
}

func (s *quickSorter) partition(a []int, lo, hi int, less func(i, j int) bool) int {
	v := a[lo]
	j := lo + 1 // [lo+1, j)<v, [j, i)>v
	for i := lo + 1; i <= hi; i++ {
		if less(a[i], v) {
			Exch(a, i, j) // [lo+1, j+1)<v, [j+1, i+1)>v
			j++           // [lo+1, j)<v, [j, i+1)>v
		} // [lo+1, j)<v, [j, i)>v
	}
	Exch(a, lo, j-1)
	return j - 1
}

func (s *quickSorter) partition2(a []int, lo, hi int, less func(i, j int) bool) int {
	Exch(a, lo, rand.Intn(hi-lo+1)+lo)
	v := a[lo]
	i, j := lo+1, hi // [lo+1, i)<=v, (j, hi]>=v
	for {
		for ; i <= hi && less(a[i], v); i++ {
		} // a[i] >=v
		for ; j >= lo+1 && less(v, a[j]); j-- {
		} // a[i] <= v
		if i >= j {
			break // [lo+1, i)<=v, (j, hi]>=v
		}
		Exch(a, i, j) // [lo+1, i] <= v, [j, hi] >=v
		i++
		j--
		// [lo+1, i)<=v, (j, hi]>=v
	}
	// i对应第一个大于等于v的元素，j对应最后一个小于等于v的元素（从后往前看第一个小于等于v的元素）
	Exch(a, lo, j)
	return j
}

// ----------------------------
// three way quick sort
// ----------------------------

type threeWayQuickSorter struct {
	*baseSorter
}

func NewThreeWayQuick() *threeWayQuickSorter {
	impl := &threeWayQuickSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *threeWayQuickSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	s.sort(a, 0, n-1, less)
}

func (s *threeWayQuickSorter) threeWayPartition1(a []int, lo, hi int, less func(i, j int) bool) (int, int) {
	Exch(a, lo, rand.Intn(hi-lo+1)+lo)
	v := a[lo]
	// [lo +1, lt) < v , [lt, i)==v, (gt, hi]>v
	lt, gt := lo+1, hi
	i := lo + 1
	for i <= gt {
		if less(a[i], v) {
			Exch(a, lt, i)
			lt++
			i++
		} else if less(v, a[i]) {
			Exch(a, gt, i)
			gt--
		} else {
			i++
		}
	}
	// i-1 == gt
	Exch(a, lo, lt-1)
	// [lo, lt-1) < v, [lt-1, gt] ==v , (gt, hi]>v
	return lt - 2, gt + 1
}

func (s *threeWayQuickSorter) threeWayPartition2(a []int, lo, hi int, less func(i, j int) bool) (int, int) {
	Exch(a, lo, rand.Intn(hi-lo+1)+lo)
	v := a[lo]
	// [lo, lt) < v , [lt, i)==v, (gt, hi]>v
	lt, gt := lo, hi
	i := lo + 1
	for i <= gt {
		if less(a[i], v) { // a[i] < v
			Exch(a, lt, i)
			lt++
			i++
		} else if a[i] == v { //a[i] ==v
			i++
		} else { // a[i] > v
			Exch(a, gt, i)
			gt--
		}
	}
	// i == gt + 1
	// [lo, lt-1] < v, [lt, gt] ==v , (gt, hi]>v
	return lt - 1, gt + 1
}

func (s *threeWayQuickSorter) BentleyMcIlroyPartition(a []int, lo, hi int, less func(i, j int) bool) (int, int) {
	Exch(a, lo, rand.Intn(hi-lo+1)+lo)
	v := a[lo]
	// [lo, p) ==v, [p, i) <v, (j, q]>v, (q, hi]==v
	p, q := lo+1, hi
	i, j := lo+1, hi
	for {
		for ; i <= hi && less(a[i], v); i++ {
		} // a[i] >= v
		for ; less(v, a[j]); j-- {
		} // a[j] <= v

		if i > j { // [lo, i) <=v , (j, hi]>=v
			break
		}
		Exch(a, i, j) // [lo, i] <=v , [j, hi]>=v
		if a[i] == v {
			Exch(a, i, p)
			p++
		}
		if a[j] == v {
			Exch(a, j, q)
			q--
		}
		i++
		j--
	}
	// move, [lo, i) <= v , (j, hi] >= v, i = j + 1
	for k := lo; k < p; k++ {
		Exch(a, k, j)
		j--
	}
	for k := hi; k > q; k-- {
		Exch(a, k, i)
		i++
	}
	return j, i
}

func (s *threeWayQuickSorter) sort(a []int, lo, hi int, less func(i, j int) bool) {
	if hi <= lo {
		return
	}

	lt, gt := s.BentleyMcIlroyPartition(a, lo, hi, less)
	s.sort(a, lo, lt, less)
	s.sort(a, gt, hi, less)
}

// ----------------------------
// advanced quick sort
// ----------------------------
type advancedQuickSorter struct {
	*baseSorter
	switchSorter abstract.Sorter
}

func NewAdvancedQuick() *advancedQuickSorter {
	impl := &advancedQuickSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	impl.switchSorter = NewInsertion()
	return impl
}

func (s *advancedQuickSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	s.sort(a, 0, n-1, less)
}

func (s *advancedQuickSorter) sort(a []int, lo, hi int, less func(i, j int) bool) {
	if hi-lo <= 18 {
		s.switchSorter.IndexSort(a[lo:hi+1], less)
		return
	}

	lt, gt := s.BentleyMcIlroyPartition(a, lo, hi, less)
	s.sort(a, lo, lt, less)
	s.sort(a, gt, hi, less)
}

func (s *advancedQuickSorter) BentleyMcIlroyPartition(a []int, lo, hi int, less func(i, j int) bool) (int, int) {
	Exch(a, lo, median3(a, lo, hi, less))
	v := a[lo]
	// [lo, p) ==v, [p, i) <v, (j, q]>v, (q, hi]==v
	p, q := lo+1, hi
	i, j := lo+1, hi
	for {
		for ; i <= hi && less(a[i], v); i++ {
		} // a[i] >= v
		for ; less(v, a[j]); j-- {
		} // a[j] <= v

		if i > j { // [lo, i) <=v , (j, hi]>=v
			break
		}
		Exch(a, i, j) // [lo, i] <=v , [j, hi]>=v
		if a[i] == v {
			Exch(a, i, p)
			p++
		}
		if a[j] == v {
			Exch(a, j, q)
			q--
		}
		i++
		j--
	}
	// move, [lo, i) <= v , (j, hi] >= v, i = j + 1
	for k := lo; k < p; k++ {
		Exch(a, k, j)
		j--
	}
	for k := hi; k > q; k-- {
		Exch(a, k, i)
		i++
	}
	return j, i
}

func median3(a []int, lo, hi int, less func(i, j int) bool) int {
	if hi-lo+1 < 3 {
		return lo
	}
	samplings := []int{
		rand.Intn(hi-lo+1) + lo,
		rand.Intn(hi-lo+1) + lo,
		rand.Intn(hi-lo+1) + lo,
	}
	// [0, i)sorted, [i, 2]unsorted
	for i := 1; i <= 2; i++ {
		for j := i; j > 0; j-- {
			if less(a[samplings[j]], a[samplings[j-1]]) {
				Exch(samplings, j, j-1)
			}
		}
	}
	return samplings[1]
}
