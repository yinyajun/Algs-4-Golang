/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/28 18:05
 */

package sorting

type selectionSorter struct {
	*baseSorter
}

func NewSelection() *selectionSorter {
	impl := &selectionSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *selectionSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	// [0, i) sorted, [i, length) unsorted
	for i := 0; i < n; i++ {
		minValueIndex := i
		for j := i + 1; j < n; j++ {
			if less(a[j], a[minValueIndex]) {
				minValueIndex = j
			}
		}
		// find min_val_idx in [i, length)
		Exch(a, i, minValueIndex)
	}
}
