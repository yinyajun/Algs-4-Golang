package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

/**

*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)
	fmt.Println(g)
}
