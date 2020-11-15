/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 10:18
 */

package sorting

import "Algs-4-Golang/abstract"

type insertingSorter struct {
	*selectionSorter
}

func NewInsertionSorter(indexer abstract.Indexer) *insertingSorter {
	return &insertingSorter{NewSelectionSorter(indexer)}
}

// [0, i) sorted, [i, lenght) to be sort
func (s *insertingSorter) Sort(a interface{}) {
	s.init(a)
	for i := 1; i < s.length; i++ {
		// 将a[i]插入到a[i-1],a[i-2]...
		// swap in pairs
		for j := i; j > 0 && s.Less(j, j-1); j-- {
			s.Exch(j, j-1)
		}
	}
}
