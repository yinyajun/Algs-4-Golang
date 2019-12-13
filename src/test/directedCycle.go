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

	finder := NewDirectedCycle(g)
	if finder.HasCycle() {
		fmt.Println("Directed cycle:")
		cyc := finder.Cycle()
		for ok, v := cyc(); ok; ok, v = cyc() {
			fmt.Print(v.(int), " ")
		}
		fmt.Println()
	} else {
		fmt.Print("No directed cycle")
	}
	fmt.Println()
}
