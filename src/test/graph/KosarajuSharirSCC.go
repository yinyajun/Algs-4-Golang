package main

import (
	"fmt"
	"os"

	. "algs4/bag"
	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/KosarajuSharirSCC.go < data/tinyDG.txt
* 5 components
* 1
* 5 4 3 2 0
* 12 11 10 9
* 8 6
* 7
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)
	scc := NewKosarajuSharirScc(g)

	// number of connected components
	fmt.Println(scc.Count(), "components")

	// compute list of vertices in each connected component
	components := make([]*Bag, scc.Count())
	for idx := range components {
		components[idx] = NewBag()
	}

	for v := 0; v < g.V(); v++ {
		components[scc.Id(v)].Add(v)
	}

	for _, component := range components {
		c := component.Iterate()
		for v := c.Next(); v != nil; v = c.Next() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
