package main

import (
	"os"
	"strconv"

	. "algs4/graph"
	"fmt"
	. "util"
)

/**
* $ go run src/test/DepthFirstDirectedPaths.go < data/tinyDG.txt 0
* 0 to 0:  0
* 0 to 1:  0-1
* 0 to 2:  0-5-4-3-2
* 0 to 3:  0-5-4-3
* 0 to 4:  0-5-4
* 0 to 5:  0-5
* 0 to 6: not connected
* 0 to 7: not connected
* 0 to 8: not connected
* 0 to 9: not connected
* 0 to 10: not connected
* 0 to 11: not connected
* 0 to 12: not connected
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
			for x := path.Next(); x != nil; x = path.Next() {
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
