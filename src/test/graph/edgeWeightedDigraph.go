package main

import (
	"fmt"
	"os"

	"algs4/graph"
	. "util"
)

/**
* $ go run src/test/edgeWeightedDigraph.go < data/tinyEWG.txt
* 8 16
* 0: 0 -> 2  0.26 0 -> 4  0.38 0 -> 7  0.16
* 1: 1 -> 3  0.29 1 -> 2  0.36 1 -> 7  0.19 1 -> 5  0.32
* 2: 2 -> 7  0.34 2 -> 3  0.17
* 3: 3 -> 6  0.52
* 4: 4 -> 7  0.37 4 -> 5  0.35
* 5: 5 -> 7  0.28
* 6: 6 -> 4  0.93 6 -> 0  0.58 6 -> 2  0.40
* 7:
*
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedDigraphWithIn(in)
	fmt.Println(g)
}
