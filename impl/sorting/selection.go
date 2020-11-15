/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:27
 */

package sorting

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
	"reflect"
)

type selectionSorter struct {
	length  int
	indexer abstract.Indexer // Lack of generic type, use this to index value in a interface(slice type)
	swapper func(i, j int)
}

func NewSelectionSorter(indexer abstract.Indexer) *selectionSorter {
	return &selectionSorter{indexer: indexer}
}

func (s *selectionSorter) init(a interface{}) {
	if s.swapper == nil {
		s.length = reflect.ValueOf(a).Len()
		s.swapper = reflect.Swapper(a)
	}
}

// left: [0, i) is sorted, [i, length) to be sort
func (s *selectionSorter) Sort(a interface{}) {
	s.init(a)
	for i := 0; i < s.length; i++ {
		min_index := i
		// find the index of min value in [i, length)
		for j := i; j < s.length; j++ {
			if s.Less(j, min_index) {
				min_index = j
			}
		}
		s.Exch(min_index, i) // n swap
	}
}

func (s *selectionSorter) Less(i, j int) bool {
	return utils.Less(s.indexer(i), s.indexer(j)) // use compartor
}

func (s *selectionSorter) Exch(i, j int) { s.swapper(i, j) }

func (s *selectionSorter) IsSorted(a interface{}) bool {
	for i := 1; i < s.length; i++ {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *selectionSorter) Show(a interface{}) {
	for i := 0; i < s.length; i++ {
		utils.StdOut.Println(s.indexer(i))
	}
}
