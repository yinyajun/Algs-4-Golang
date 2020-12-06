/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:44
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/sorting"
	"Algs-4-Golang/utils"
)

// go run impl/sorting/examples/sorter.go Selection  < data/tiny.txt
// +-----------+
// | Selection |
// +-----------+
// A
// E
// E
// L
// M
// O
// P
// R
// S
// T
// X

var s abstract.Sorter

const (
	Selection         = "Selection"
	Insertion         = "Insertion"
	AdvancedInsertion = "AdvancedInsertion"
	Shell             = "Shell"
	Merge             = "Merge"
	AdvancedMerge     = "AdvancedMerge"
	Quick             = "Quick"
	ThreeWayQuick     = "ThreeWayQuick"
	AdvancedQuick     = "AdvancedQuick"
	Heap              = "Heap"
)

func init() {
	utils.Arg0 = utils.Flag.Arg(0, Selection)
	utils.PrintInBox(utils.Arg0)
}

func initSorter(args ...interface{}) {
	typ := args[0]
	switch typ {
	case Selection:
		s = sorting.NewSelection()
	case Insertion:
		s = sorting.NewInsertion()
	case AdvancedInsertion:
		s = sorting.NewAdvancedInsertion()
	case Shell:
		s = sorting.NewShell()
	case Merge:
		s = sorting.NewMerge()
	case AdvancedMerge:
		s = sorting.NewAdvancedMerge()
	case Quick:
		s = sorting.NewQuick()
	case ThreeWayQuick:
		s = sorting.NewThreeWayQuick()
	case AdvancedQuick:
		s = sorting.NewAdvancedQuick()
	case Heap:
		s = sorting.NewHeap()
	default:
		utils.Panic("unsupported type")
	}
}

func main() {
	a := utils.StdIn.ReadAllStrings()
	initSorter(utils.Arg0)
	s.Sort(a, func(i, j int) bool { return utils.Less(a[i], a[j]) })
	utils.Assert(s.IsSorted(a, func(i, j int) bool { return utils.Less(a[i], a[j]) }))
	s.Show(a)
}
