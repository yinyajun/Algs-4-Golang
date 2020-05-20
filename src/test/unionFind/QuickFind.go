package main

import (
	"os"
	"fmt"

	. "algs4/unionFind"
	. "util"
)

/**
*
$ go run src/test/unionFind/QuickFind.go  < data/tinyUF.txt
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

func main() {
	in := NewIn(os.Stdin)
	N := in.ReadInt()
	uf := NewQuickFindUF(N)
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
