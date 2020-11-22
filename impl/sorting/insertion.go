/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 10:18
 */

package sorting

import "reflect"

type insertingSorter struct {
	*selectionSorter
}

func NewInsertion() *insertingSorter {
	return &insertingSorter{NewSelection()}
}

// [0, i) sorted, [i, length) to be sort
func (s *insertingSorter) Sort(a interface{}, less func(i, j int) bool) {
	s.length = reflect.ValueOf(a).Len()
	s.swapper = reflect.Swapper(a)
	for i := 1; i < s.length; i++ {
		// 将a[i]插入到a[i-1],a[i-2]...
		// swap in pairs
		for j := i; j > 0 && less(j, j-1); j-- {
			s.Exch(j, j-1)
		}
	}
}
