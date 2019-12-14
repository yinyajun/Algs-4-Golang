package main

import (
	"fmt"
	"os"
	"strconv"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/breadthFirstPaths.go < data/tinyCG.txt 0
* 0 to 0 (0): 0
* 0 to 1 (1): 0-1
* 0 to 2 (1): 0-2
* 0 to 3 (2): 0-2-3
* 0 to 4 (2): 0-2-4
* 0 to 5 (1): 0-5
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])
	bfs := NewBreadthFirstPaths(g, s)
	for v := 0; v < g.V(); v++ {
		if bfs.HasPathTo(v) {
			fmt.Print(s, " to ", v, " (", bfs.DistTo(v), "): ")
			path := bfs.PathTo(v)
			for x := path.Next(); x != nil; x = path.Next() {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Print("-", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Print(s, " to ", v, " (-): ", "not connected\n")
		}
	}
}
