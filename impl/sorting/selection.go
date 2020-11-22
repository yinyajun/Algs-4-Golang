/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:27
 */

package sorting

import (
	"reflect"
)

type selectionSorter struct {
	length  int
	swapper func(i, j int)
}

func NewSelection() *selectionSorter { return &selectionSorter{} }

func (s *selectionSorter) Sort(a interface{}, less func(i, j int) bool) {
	s.length = reflect.ValueOf(a).Len()
	s.swapper = reflect.Swapper(a)
	// [0, i) sorted, [i, length) unsorted
	for i := 0; i < s.length; i++ {
		min_val_idx := i
		for j := i; j < s.length; j++ {
			if less(j, min_val_idx) {
				min_val_idx = j
			}
		}
		// find min_val_idx in [i, length)
		s.Exch(i, min_val_idx)
	}
}

func (s *selectionSorter) Exch(i, j int) { s.swapper(i, j) }

func (s *selectionSorter) IsSorted(a interface{}, less func(i, j int) bool) bool {
	for i := 1; i < s.length; i++ {
		if less(i, i-1) {
			return false
		}
	}
	return true
}
