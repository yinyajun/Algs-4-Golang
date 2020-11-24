/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/24 11:56
 */

package sorting

import "reflect"

type shellSorter struct {
	*selectionSorter
}

func NewShell() *shellSorter {
	return &shellSorter{NewSelection()}
}

func (s *shellSorter) Sort(a interface{}, less func(i, j int) bool) {
	sliceValue := reflect.ValueOf(a)
	swapper := reflect.Swapper(a)
	N := sliceValue.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && less(j, j-h); j -= h {
				swapper(j, j-h)
			}
		}
		h /= 3
	}
}
