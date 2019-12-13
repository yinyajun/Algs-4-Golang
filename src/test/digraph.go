package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/digraph.go < data/tinyDG.txt
* 13 vertices, 22 edges
* 0: 5 1
* 1:
* 2: 0 3
* 3: 5 2
* 4: 3 2
* 5: 4
* 6: 9 4 8 0
* 7: 6 9
* 8: 6
* 9: 11 10
* 10: 12
* 11: 4 12
* 12: 9
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)
	fmt.Println(g)
}
