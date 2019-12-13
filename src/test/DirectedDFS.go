package main

import (
	"fmt"
	"os"
	"strconv"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/depthFirstSearch.go < data/tinyCG.txt 0
* 0 1 2 3 4 5
* connected
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)

	// read in sources
	sources := []int{}
	for i := 1; i < len(os.Args); i++ {
		s, _ := strconv.Atoi(os.Args[i])
		sources = append(sources, s)
	}

	// multiple-source reachability
	dfs := NewDirectedDFSMultiSources(g, sources)

	// print out vertices reachable from sources
	for v := 0; v < g.V(); v++ {
		if dfs.Marked(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
