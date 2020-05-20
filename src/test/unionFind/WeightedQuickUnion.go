package main

/**
*
$ go run src/test/unionFind/WeightedQuickUnion.go  < data/tinyUF.txt
4 3
3 8
6 5
9 4
2 1
5 0
7 2
6 1
2 components
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

import (
	"os"
	"fmt"

	. "algs4/unionFind"
	. "util"
)

func main() {
	in := NewIn(os.Stdin)
	N := in.ReadInt()
	uf := NewWeightedQuickUnionUF(N)
	for in.HasNext() {
		p := in.ReadInt()
		q := in.ReadInt()
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(uf.Count(), "components")
}
