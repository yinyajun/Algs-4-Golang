/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 11:37
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

// go run impl/fundamentals/examples/union_find.go  < data/tinyUF.txt
// +-------------+
// | QuickFindUF |
// +-------------+
// 4 3
// 3 8
// 6 5
// 9 4
// 2 1
// 5 0
// 7 2
// 6 1
// 2 components

var uf abstract.UnionFind

const (
	QuickFindUF             = "QuickFindUF"
	QuickUnionUF            = "QuickUnionUF"
	QuickUnionSizeUF        = "QuickUnionSizeUF"
	QuickUnionRankUF        = "QuickUnionRankUF"
	QuickUnionCompressedUF  = "QuickUnionCompressedUF"
	QuickUnionCompressedUF2 = "QuickUnionCompressedUF2"
)

func init() {
	utils.Arg0 = utils.Flag.Arg(0, QuickFindUF)
	utils.PrintInBox(utils.Arg0)
}

func initUnionFind(args ...interface{}) {
	typ := args[0]
	n := args[1].(int)
	switch typ {
	case QuickFindUF:
		uf = fundamentals.NewQuickFindUF(n)
	case QuickUnionUF:
		uf = fundamentals.NewQuickUnionUF(n)
	case QuickUnionSizeUF:
		uf = fundamentals.NewQuickUnionSizeUF(n)
	case QuickUnionRankUF:
		uf = fundamentals.NewQuickUnionRankUF(n)
	case QuickUnionCompressedUF:
		uf = fundamentals.NewQuickUnionCompressed(n)
	case QuickUnionCompressedUF2:
		uf = fundamentals.NewQuickUnionCompressed2(n)
	}
}

func main() {
	N := utils.StdIn.ReadInt()
	initUnionFind(utils.Arg0, N)
	for utils.StdIn.HasNext() {
		p := utils.StdIn.ReadInt()
		q := utils.StdIn.ReadInt()
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
		utils.StdOut.Println(p, q)
	}
	utils.StdOut.Println(uf.Count(), "components")
}
