package main

import (
	"fmt"
	"os"

	. "util"
	"algs4/graph"
)

/**
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedGraphWithIn(in)
	mst := graph.NewLazyPrimMST(g)

	edges := mst.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f\n", mst.Weight())
}
