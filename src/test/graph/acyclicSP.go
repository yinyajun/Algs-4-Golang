package main

import (
	"fmt"
	"os"
	"strconv"

	"algs4/graph"
	. "util"
)

/**
* $ go run src/test/acyclicSP.go 5 < data/tinyEWDAG.txt
* 5 to 0 (0.73)  5 -> 4  0.35   4 -> 0  0.38
* 5 to 1 (0.32)  5 -> 1  0.32
* 5 to 2 (0.62)  5 -> 7  0.28   7 -> 2  0.34
* 5 to 3 (0.61)  5 -> 1  0.32   1 -> 3  0.29
* 5 to 4 (0.35)  5 -> 4  0.35
* 5 to 5 (0.00)
* 5 to 6 (1.13)  5 -> 1  0.32   1 -> 3  0.29   3 -> 6  0.52
* 5 to 7 (0.28)  5 -> 7  0.28
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedDigraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])

	// find shortest path from s to each other vertex in DAG
	sp := graph.NewAcyclicSP(g, s)
	for v := 0; v < g.V(); v ++ {
		if sp.HasPathTo(v) {
			fmt.Printf("%d to %d (%.2f)  ", s, v, sp.DistTo(v))
			edges := sp.PathTo(v)
			for e := edges.Next(); e != nil; e = edges.Next() {
				fmt.Print(e, "   ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, v)
		}
	}
}
