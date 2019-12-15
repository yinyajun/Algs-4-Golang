package main

import (
	"fmt"
	"os"

	. "util"
	"algs4/graph"
)

/**
* $ go run src/test/edgeWeightedGraph.go < data/tinyEWG.txt
* 8 16
* 0 : 6-0 0.58000 0-2 0.26000 0-4 0.38000 0-7 0.16000
* 1 : 1-3 0.29000 1-2 0.36000 1-7 0.19000 1-5 0.32000
* 2 : 6-2 0.40000 2-7 0.34000 1-2 0.36000 0-2 0.26000 2-3 0.17000
* 3 : 3-6 0.52000 1-3 0.29000 2-3 0.17000
* 4 : 6-4 0.93000 0-4 0.38000 4-7 0.37000 4-5 0.35000
* 5 : 1-5 0.32000 5-7 0.28000 4-5 0.35000
* 6 : 6-4 0.93000 6-0 0.58000 3-6 0.52000 6-2 0.40000
* 7 : 2-7 0.34000 1-7 0.19000 0-7 0.16000 5-7 0.28000 4-7 0.37000
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := graph.NewEdgeWeightedGraphWithIn(in)
	fmt.Println(g)
}
