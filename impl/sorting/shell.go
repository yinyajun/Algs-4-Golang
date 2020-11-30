/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/29 11:26
 */

package sorting

// ----------------------------
// shell sort
// ----------------------------

type shellSorter struct {
	*baseSorter
}

func NewShell() *shellSorter {
	impl := &shellSorter{}
	base := &baseSorter{}
	base.impl = impl
	impl.baseSorter = base
	return impl
}

func (s *shellSorter) IndexSort(a []int, less func(i, j int) bool) {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && less(a[j], a[j-h]); j -= h {
				Exch(a, j, j-h)
			}
		}
		h /= 3
	}
}
