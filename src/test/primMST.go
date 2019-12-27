package main

import (
	"fmt"
	"os"

	"algs4/graph"
	. "util"
)

/**
*

*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedGraphWithIn(in)
	mst := graph.NewPrimMST(g)

	edges := mst.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f\n", mst.Weight())
}
