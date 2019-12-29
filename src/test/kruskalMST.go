package main

import (
	"fmt"
	"os"

	"algs4/graph"
	. "util"
)

/**
* go run src/test/kruskalMST.go < data/tinyEWG.txt
* 0-7 0.16000
* 2-3 0.17000
* 1-7 0.19000
* 0-2 0.26000
* 5-7 0.28000
* 4-5 0.35000
* 6-2 0.40000
* 1.81000
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedGraphWithIn(in)
	mst := graph.NewKruskalMST(g)

	edges := mst.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f\n", mst.Weight())
}
