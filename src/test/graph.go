package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

/**
*  $ go run src/test/graph.go < data/tinyG.txt
*  13 vertices, 13 edges
*  0:6 2 1 5
*  1:0
*  2:0
*  3:5 4
*  4:5 6 3
*  5:3 4 0
*  6:0 4
*  7:8
*  8:7
*  9:11 10 12
*  10:9
*  11:9 12
*  12:11 9
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	fmt.Println(g)
}
