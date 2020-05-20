package main

import (
	"fmt"
	"os"
	"strconv"

	"algs4/graph"
	. "util"
)

/**
* $ go run src/test/dijkstra.go 0 < data/tinyEWD.txt
* 0 to 0 (0.00)
* 0 to 1 (1.05)  0 -> 4  0.38   4 -> 5  0.35   5 -> 1  0.32
* 0 to 2 (0.26)  0 -> 2  0.26
* 0 to 3 (0.99)  0 -> 2  0.26   2 -> 7  0.34   7 -> 3  0.39
* 0 to 4 (0.38)  0 -> 4  0.38
* 0 to 5 (0.73)  0 -> 4  0.38   4 -> 5  0.35
* 0 to 6 (1.51)  0 -> 2  0.26   2 -> 7  0.34   7 -> 3  0.39   3 -> 6  0.52
* 0 to 7 (0.60)  0 -> 2  0.26   2 -> 7  0.34
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedDigraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])

	// compute shortest paths
	sp := graph.NewDijkstraSP(g, s)

	// print shortest path
	for t := 0; t < g.V(); t++ {
		if sp.HasPathTo(t) {
			fmt.Printf("%d to %d (%.2f)  ", s, t, sp.DistTo(t))
			edges := sp.PathTo(t)
			for e := edges.Next(); e != nil; e = edges.Next() {
				fmt.Print(e, "   ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, t)
		}
	}
}
