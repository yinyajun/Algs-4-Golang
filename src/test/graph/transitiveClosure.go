package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/transitiveClosure.go < data/tinyDG.txt
*        0  1  2  3  4  5  6  7  8  9 10 11 12
* --------------------------------------------
*   0  T  T  T  T  T  T
*   1     T
*   2  T  T  T  T  T  T
*   3  T  T  T  T  T  T
*   4  T  T  T  T  T  T
*   5  T  T  T  T  T  T
*   6  T  T  T  T  T  T  T     T  T  T  T  T
*   7  T  T  T  T  T  T  T  T  T  T  T  T  T
*   8  T  T  T  T  T  T  T     T  T  T  T  T
*   9  T  T  T  T  T  T           T  T  T  T
*  10  T  T  T  T  T  T           T  T  T  T
*  11  T  T  T  T  T  T           T  T  T  T
*  12  T  T  T  T  T  T           T  T  T  T
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)
	tc := NewTransitiveClosure(g)

	fmt.Print("     ")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("%3d", v)
	}
	fmt.Println()
	fmt.Println("--------------------------------------------")

	for v := 0; v < g.V(); v++ {
		fmt.Printf("%3d", v)
		for w := 0; w < g.V(); w++ {
			if tc.Reachable(v, w) {
				fmt.Print("  T")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}
