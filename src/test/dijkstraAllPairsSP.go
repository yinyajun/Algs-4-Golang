package main

import (
	"algs4/graph"
	"fmt"
	"os"
	. "util"
)

/**
* $ go run src/test/dijkstraAllPairsSP.go < data/tinyEWD.txt
*        0      1      2      3      4      5      6      7
*   0:   0.00   1.05   0.26   0.99   0.38   0.73   1.51   0.60
*   1:   1.39   0.00   1.21   0.29   1.74   1.83   0.81   1.55
*   2:   1.83   0.94   0.00   0.73   0.97   0.62   1.25   0.34
*   3:   1.10   1.86   0.92   0.00   1.45   1.54   0.52   1.26
*   4:   1.86   0.67   1.68   0.76   0.00   0.35   1.28   0.37
*   5:   1.71   0.32   1.53   0.61   0.35   0.00   1.13   0.28
*   6:   0.58   1.34   0.40   1.13   0.93   1.02   0.00   0.74
*   7:   1.49   0.60   1.31   0.39   0.63   0.28   0.91   0.00
*
* 0 to 0 ( 0.00)
* 0 to 1 ( 1.05)  0 -> 4  0.38  4 -> 5  0.35  5 -> 1  0.32
* 0 to 2 ( 0.26)  0 -> 2  0.26
* 0 to 3 ( 0.99)  0 -> 2  0.26  2 -> 7  0.34  7 -> 3  0.39
* 0 to 4 ( 0.38)  0 -> 4  0.38
* 0 to 5 ( 0.73)  0 -> 4  0.38  4 -> 5  0.35
* 0 to 6 ( 1.51)  0 -> 2  0.26  2 -> 7  0.34  7 -> 3  0.39  3 -> 6  0.52
* 0 to 7 ( 0.60)  0 -> 2  0.26  2 -> 7  0.34
* 1 to 0 ( 1.39)  1 -> 3  0.29  3 -> 6  0.52  6 -> 0  0.58
* 1 to 1 ( 0.00)
* 1 to 2 ( 1.21)  1 -> 3  0.29  3 -> 6  0.52  6 -> 2  0.40
* 1 to 3 ( 0.29)  1 -> 3  0.29
* 1 to 4 ( 1.74)  1 -> 3  0.29  3 -> 6  0.52  6 -> 4  0.93
* 1 to 5 ( 1.83)  1 -> 3  0.29  3 -> 6  0.52  6 -> 2  0.40  2 -> 7  0.34  7 -> 5  0.28
* 1 to 6 ( 0.81)  1 -> 3  0.29  3 -> 6  0.52
* 1 to 7 ( 1.55)  1 -> 3  0.29  3 -> 6  0.52  6 -> 2  0.40  2 -> 7  0.34
* 2 to 0 ( 1.83)  2 -> 7  0.34  7 -> 3  0.39  3 -> 6  0.52  6 -> 0  0.58
* 2 to 1 ( 0.94)  2 -> 7  0.34  7 -> 5  0.28  5 -> 1  0.32
* 2 to 2 ( 0.00)
* 2 to 3 ( 0.73)  2 -> 7  0.34  7 -> 3  0.39
* 2 to 4 ( 0.97)  2 -> 7  0.34  7 -> 5  0.28  5 -> 4  0.35
* 2 to 5 ( 0.62)  2 -> 7  0.34  7 -> 5  0.28
* 2 to 6 ( 1.25)  2 -> 7  0.34  7 -> 3  0.39  3 -> 6  0.52
* 2 to 7 ( 0.34)  2 -> 7  0.34
* 3 to 0 ( 1.10)  3 -> 6  0.52  6 -> 0  0.58
* 3 to 1 ( 1.86)  3 -> 6  0.52  6 -> 2  0.40  2 -> 7  0.34  7 -> 5  0.28  5 -> 1  0.32
* 3 to 2 ( 0.92)  3 -> 6  0.52  6 -> 2  0.40
* 3 to 3 ( 0.00)
* 3 to 4 ( 1.45)  3 -> 6  0.52  6 -> 4  0.93
* 3 to 5 ( 1.54)  3 -> 6  0.52  6 -> 2  0.40  2 -> 7  0.34  7 -> 5  0.28
* 3 to 6 ( 0.52)  3 -> 6  0.52
* 3 to 7 ( 1.26)  3 -> 6  0.52  6 -> 2  0.40  2 -> 7  0.34
* 4 to 0 ( 1.86)  4 -> 7  0.37  7 -> 3  0.39  3 -> 6  0.52  6 -> 0  0.58
* 4 to 1 ( 0.67)  4 -> 5  0.35  5 -> 1  0.32
* 4 to 2 ( 1.68)  4 -> 7  0.37  7 -> 3  0.39  3 -> 6  0.52  6 -> 2  0.40
* 4 to 3 ( 0.76)  4 -> 7  0.37  7 -> 3  0.39
* 4 to 4 ( 0.00)
* 4 to 5 ( 0.35)  4 -> 5  0.35
* 4 to 6 ( 1.28)  4 -> 7  0.37  7 -> 3  0.39  3 -> 6  0.52
* 4 to 7 ( 0.37)  4 -> 7  0.37
* 5 to 0 ( 1.71)  5 -> 1  0.32  1 -> 3  0.29  3 -> 6  0.52  6 -> 0  0.58
* 5 to 1 ( 0.32)  5 -> 1  0.32
* 5 to 2 ( 1.53)  5 -> 1  0.32  1 -> 3  0.29  3 -> 6  0.52  6 -> 2  0.40
* 5 to 3 ( 0.61)  5 -> 1  0.32  1 -> 3  0.29
* 5 to 4 ( 0.35)  5 -> 4  0.35
* 5 to 5 ( 0.00)
* 5 to 6 ( 1.13)  5 -> 1  0.32  1 -> 3  0.29  3 -> 6  0.52
* 5 to 7 ( 0.28)  5 -> 7  0.28
* 6 to 0 ( 0.58)  6 -> 0  0.58
* 6 to 1 ( 1.34)  6 -> 2  0.40  2 -> 7  0.34  7 -> 5  0.28  5 -> 1  0.32
* 6 to 2 ( 0.40)  6 -> 2  0.40
* 6 to 3 ( 1.13)  6 -> 2  0.40  2 -> 7  0.34  7 -> 3  0.39
* 6 to 4 ( 0.93)  6 -> 4  0.93
* 6 to 5 ( 1.02)  6 -> 2  0.40  2 -> 7  0.34  7 -> 5  0.28
* 6 to 6 ( 0.00)
* 6 to 7 ( 0.74)  6 -> 2  0.40  2 -> 7  0.34
* 7 to 0 ( 1.49)  7 -> 3  0.39  3 -> 6  0.52  6 -> 0  0.58
* 7 to 1 ( 0.60)  7 -> 5  0.28  5 -> 1  0.32
* 7 to 2 ( 1.31)  7 -> 3  0.39  3 -> 6  0.52  6 -> 2  0.40
* 7 to 3 ( 0.39)  7 -> 3  0.39
* 7 to 4 ( 0.63)  7 -> 5  0.28  5 -> 4  0.35
* 7 to 5 ( 0.28)  7 -> 5  0.28
* 7 to 6 ( 0.91)  7 -> 3  0.39  3 -> 6  0.52
* 7 to 7 ( 0.00)
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	// read edge-weighted digraph
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedDigraphWithIn(in)

	// compute shortest paths between all pairs of vertices
	spt := graph.NewDijkstraAllPairsSP(g)

	// print all-pairs shortest path distances
	fmt.Print("  ")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("%6d ", v)
	}
	fmt.Println()
	for v := 0; v < g.V(); v++ {
		fmt.Printf("%3d: ", v)
		for w := 0; w < g.V(); w++ {
			if spt.HasPath(v, w) {
				fmt.Printf("%6.2f ", spt.Dist(v, w))
			} else {
				fmt.Print("  Inf ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// print all-pairs shortest paths
	for v := 0; v < g.V(); v++ {
		for w := 0; w < g.V(); w++ {
			if spt.HasPath(v, w) {
				fmt.Printf("%d to %d (%5.2f)  ", v, w, spt.Dist(v, w))
				path := spt.Path(v, w)
				for e := path.Next(); e != nil; e = path.Next() {
					fmt.Print(e)
					fmt.Print("  ")
				}
				fmt.Println()
			} else {
				fmt.Printf("%d to %d no path\n", v, w)
			}
		}
	}
}
