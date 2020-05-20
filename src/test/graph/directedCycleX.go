package main

import (
	"os"
)

import (
	. "algs4/graph"
	"fmt"
	. "util"
)

/**
* $ go run src/test/directedCycleX.go < data/tinyDG.txt
* Directed cycle:
* 11 12 9 11
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewDigraphWithIn(in)

	finder := NewDirectedCycleX(g)
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
