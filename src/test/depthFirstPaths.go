package main

import (
	"fmt"
	"os"
	"strconv"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/depthFirstPaths.go < data/tinyCG.txt 0
* 0 to 0: 0
* 0 to 1: 0-2-1
* 0 to 2: 0-2
* 0 to 3: 0-2-3
* 0 to 4: 0-2-3-4
* 0 to 5: 0-2-3-5
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])
	search := NewDepthFirstPaths(g, s)
	for v := 0; v < g.V(); v++ {
		fmt.Print(s, " to ", v, ": ")
		if search.HasPathTo(v) {
			gen := search.PathTo(v)
			for hasNext, x := gen(); hasNext; hasNext, x = gen() {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Print("-", x)
				}
			}
		}
		fmt.Println()
	}
}
