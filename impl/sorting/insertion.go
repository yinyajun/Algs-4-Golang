/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/29 10:58
 */

package sorting

type insertingSorter struct {
	*baseSorter
}

func NewInsertion() *insertingSorter {
	impl := &insertingSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *insertingSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	// [0, i) sorted, [i, length) to be sort
	for i := 1; i < n; i++ {
		// 将a[i]插入到a[i-1],a[i-2]...
		for j := i; j > 0 && less(a[j], a[j-1]); j-- {
			Exch(a, j, j-1)
		}
	}
}

type advancedInsertingSorter struct {
	*baseSorter
}

func NewAdvancedInsertion() *advancedInsertingSorter {
	impl := &advancedInsertingSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

// 用赋值代替交换
func (s *advancedInsertingSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	var j int
	// [0, i) sorted, [i, length) to be sort
	for i := 1; i < n; i++ {
		e := a[i]
		for j = i; j > 0 && less(e, a[j-1]); j-- {
			a[j] = a[j-1]
		}
		// j==0 || a[j-1]<=e
		a[j] = e
	}
}
