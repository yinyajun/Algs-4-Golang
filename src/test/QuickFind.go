package main

import (
	"os"
	"fmt"

	. "algs4/unionFind"
	. "util"
)

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
