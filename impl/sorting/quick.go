/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/30 10:04
 */

package sorting

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
			Exch(a, i, j)
			j++
		}
	}
	Exch(a, lo, j-1)
	return j - 1
}

func (s *quickSorter) partition2(a []int, lo, hi int, less func(i, j int) bool) int {
	v := a[lo]
	i, j := lo+1, hi // [lo+1, i)<=v, (j, hi]>=v
	for {
		for ; i <= hi && less(a[i], v); i++ {
		} // [lo+1, i)<=v
		for ; j <= hi && less(v, a[j]); j-- {
		} // (j, hi]>=v
		if i > j {
			break
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
