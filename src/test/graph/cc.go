package main

import (
	"fmt"
	"os"

	. "algs4/bag"
	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/cc.go < data/tinyG.txt
* 3 components
* 6 5 4 3 2 1 0
* 8 7
* 12 11 10 9
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	cc := NewCC(g)

	// number of connected components
	fmt.Println(cc.Count(), "components")

	// compute list of vertices in each connected component
	components := make([]*Bag, cc.Count())
	for idx := range components {
		components[idx] = NewBag()
	}

	for v := 0; v < g.V(); v++ {
		components[cc.Id(v)].Add(v)
	}

	for _, component := range components {
		c := component.Iterate()
		for v := c.Next(); v != nil; v = c.Next() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}
