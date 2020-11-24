/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:27
 */

package sorting

import (
	"Algs-4-Golang/utils"
	"reflect"
)

type selectionSorter struct{}

func NewSelection() *selectionSorter { return &selectionSorter{} }

func (s *selectionSorter) Sort(a interface{}, less func(i, j int) bool) {
	sliceValue := reflect.ValueOf(a)
	swapper := reflect.Swapper(a)
	// [0, i) sorted, [i, length) unsorted
	for i := 0; i < sliceValue.Len(); i++ {
		min_val_idx := i
		for j := i + 1; j < sliceValue.Len(); j++ {
			if less(j, min_val_idx) {
				min_val_idx = j
			}
		}
		// find min_val_idx in [i, length)
		swapper(i, min_val_idx)
	}
}

func (s *selectionSorter) IsSorted(a interface{}, less func(i, j int) bool) bool {
	sliceValue := reflect.ValueOf(a)
	for i := 1; i < sliceValue.Len(); i++ {
		if less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *selectionSorter) Show(a interface{}) {
	sliceValue := reflect.ValueOf(a)
	for i := 0; i < sliceValue.Len(); i++ {
		utils.StdOut.Println(sliceValue.Index(i))
	}
}
