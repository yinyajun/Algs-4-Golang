/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 10:18
 */

package sorting

import (
	"Algs-4-Golang/utils"
	"reflect"
)

type insertingSorter struct {
	*selectionSorter
}

func NewInsertion() *insertingSorter {
	return &insertingSorter{NewSelection()}
}

// [0, i) sorted, [i, length) to be sort
func (s *insertingSorter) Sort(a interface{}, less func(i, j int) bool) {
	sliceValue := reflect.ValueOf(a)
	swapper := reflect.Swapper(a)
	for i := 1; i < sliceValue.Len(); i++ {
		// 将a[i]插入到a[i-1],a[i-2]...
		// swap in pairs
		for j := i; j > 0 && less(j, j-1); j-- {
			swapper(j, j-1)
		}
	}
}

type advancedInsertingSorter struct {
	*selectionSorter
}

func NewAdvancedInsertionSorter() *advancedInsertingSorter {
	return &advancedInsertingSorter{NewSelection()}
}

// [0, i) sorted, [i, length) to be sort
// reflect will make performance worse
func (s *advancedInsertingSorter) Sort(a interface{}, less func(i, j int) bool) {
	sliceValue := reflect.ValueOf(a)
	getter := func(idx int) interface{} { return sliceValue.Index(idx).Interface() }
	setter := func(idx int, v interface{}) { sliceValue.Index(idx).Set(reflect.ValueOf(v)) }

	for i := 1; i < sliceValue.Len(); i++ {
		// 将a[i]之前的较大元素都往后移
		e := getter(i)
		var j int
		for j = i; j > 0 && utils.Less(e, getter(j-1)); j-- {
			setter(j, getter(j-1))
		}
		// j == 0 || a[j-1] <= e
		setter(j, e)
	}
}
