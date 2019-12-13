package main

import (
	"os"
	"strconv"

	. "algs4/graph"
	. "util"
	"fmt"
)

/**

*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)

	s, _ := strconv.Atoi(os.Args[1])
	dfs := NewDepthFirstDirectedPaths(g, s)

	for v := 0; v < g.V(); v++ {
		if dfs.HasPathTo(v) {
			fmt.Printf("%d to %d:  ", s, v)
			path := dfs.PathTo(v)
			for ok, x := path(); ok; ok, x = path() {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Print("-", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d: not connected\n", s, v)
		}

	}
}
