/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/28 17:49
 */

package sorting

import (
	"reflect"

	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
)

type baseSorter struct {
	impl abstract.Sorter
}

func (s *baseSorter) Sort(a interface{}, less func(i, j int) bool) {
	indexes := utils.BuildIndexSlice(reflect.ValueOf(a).Len())

	s.impl.IndexSort(indexes, less)

	utils.SortByIndex(indexes, reflect.Swapper(a))
}

func (s *baseSorter) IsSorted(a interface{}, less func(i, j int) bool) bool {
	sliceValue := reflect.ValueOf(a)
	for i := 1; i < sliceValue.Len(); i++ {
		if less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *baseSorter) Show(a interface{}) {
	sliceValue := reflect.ValueOf(a)
	for i := 0; i < sliceValue.Len(); i++ {
		utils.StdOut.Println(sliceValue.Index(i))
	}
}

func Exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
