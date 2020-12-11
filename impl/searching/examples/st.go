/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/10 21:23
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/searching"
	"Algs-4-Golang/utils"
)

// go run impl/searching/examples/st.go BinarySearchST < data/tinyST.txt
// +----------------+
// | BinarySearchST |
// +----------------+
// A 8
// C 4
// E 12
// H 5
// L 11
// M 9
// P 10
// R 3
// S 0
// X 7

const (
	SequentialSearchST = "SequentialSearchST"
	BinarySearchST     = "BinarySearchST"
)

var st abstract.SymbolTable

func init() {
	utils.Arg0 = utils.Flag.Arg(0, BinarySearchST)
	utils.PrintInBox(utils.Arg0)
}

func initST(args ...interface{}) {
	typ := args[0]
	switch typ {
	case SequentialSearchST:
		st = searching.NewSequentialSearchST()
	case BinarySearchST:
		st = searching.NewBinarySearchST()
	default:
		utils.Panic("unsupported type")
	}
}

func main() {
	initST(utils.Arg0)
	for i := 0; utils.StdIn.HasNext(); i++ {
		key := utils.StdIn.ReadString()
		st.Put(key, i)
	}
	keys := st.Keys()
	for keys.First(); keys.HasNext(); {
		key := keys.Next()
		utils.StdOut.Println(key, st.Get(key))
	}
}
