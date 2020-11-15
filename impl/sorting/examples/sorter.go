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
	Selection = "Selection"
	Insertion = "Insertion"
)

func init() {
	utils.Arg0 = utils.Flag.Arg(0, Selection)
	utils.PrintInBox(utils.Arg0)
}

func initSorter(args ...interface{}) {
	typ := args[0]
	slice := args[1].([]string)
	indexer := func(i int) interface{} { return slice[i] }
	switch typ {
	case Selection:
		s = sorting.NewSelectionSorter(indexer)
	case Insertion:
		s = sorting.NewInsertionSorter(indexer)
	}
}

func main() {
	a := utils.StdIn.ReadAllStrings()
	initSorter(utils.Arg0, a)
	s.Sort(a)
	utils.Assert(s.IsSorted(a))
	s.Show(a)
}
