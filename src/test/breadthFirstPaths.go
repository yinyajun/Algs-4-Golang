package main

import (
	"fmt"
	"os"
	"strconv"

	. "algs4/graph"
	. "util"
)

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])
	bfs := NewBreadthFirstPaths(g, s)
	for v := 0; v < g.V(); v++ {
		if bfs.HasPathTo(v) {
			fmt.Print(s, " to ", v, " (",bfs.DistTo(v),"): ")
			gen := bfs.PathTo(v)
			for hasNext, x := gen(); hasNext; hasNext, x = gen() {
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
