package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

/**
* $ go run src/test/directedCycle.go < data/tinyDG.txt
* Directed cycle:
* 3 5 4 3
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)

	finder := NewDirectedCycle(g)
	if finder.HasCycle() {
		fmt.Println("Directed cycle:")
		cyc := finder.Cycle()
		for v := cyc.Next(); v != nil; v = cyc.Next() {
			fmt.Print(v.(int), " ")
		}
		fmt.Println()
	} else {
		fmt.Print("No directed cycle")
	}
	fmt.Println()
}
