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
	search := NewDFS(g, s)
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
