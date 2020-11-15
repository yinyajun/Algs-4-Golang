/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

// go run impl/fundamentals/examples/bag.go LinkedBag  < data/tobe.txt
// +-----------+
// | LinkedBag |
// +-----------+
// size of bag =  14
// is
// -
// -
// -
// that
// -
// -
// be
// -
// to
// not
// or
// be
// to

var b abstract.Bag

const (
	ResizingArrayBag = "ResizingArrayBag"
	LinkedBag        = "LinkedBag"
)

func init() {
	utils.Arg0 = utils.Flag.Arg(0, LinkedBag)
	utils.PrintInBox(utils.Arg0)
}

func initBag(args ...interface{}) {
	typ := args[0]
	switch typ {
	case ResizingArrayBag:
		b = fundamentals.NewResizingArrayBag()
	case LinkedBag:
		b = fundamentals.NewLinkedBag()
	}
}

func main() {
	initBag(utils.Arg0)
	for utils.StdIn.HasNext() {
		b.Add(utils.StdIn.ReadString())
	}
	utils.StdOut.Println("size of bag = ", b.Size())
	utils.PrintIterator(b.Iterate())
}
