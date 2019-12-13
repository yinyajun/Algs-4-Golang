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
	g := NewGraphWithIn(in)
	s, _ := strconv.Atoi(os.Args[1])
	search := NewDepthFirstSearch(g, s)
	for v := 0; v < g.V(); v++ {
		if search.Marked(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
	if search.Count() != g.V() {
		fmt.Println("Not connected")
	} else {
		fmt.Println("connected")
	}

}
