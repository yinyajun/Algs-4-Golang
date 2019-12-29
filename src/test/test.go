package main

import (
	"algs4/graph"
	"algs4/priorityQueue"
	"fmt"
	"os"
	. "util"
)

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedGraphWithIn(in)

	pq := priorityQueue.NewMaxPQwithArray(g.EdgesArray())
	fmt.Println("(", pq.Size(), "left on minPQ)")
	PrintIterator(pq.Iterate())

}
