package main

import (
	. "algs4/graph"
	"fmt"
	"os"
	. "util"
)

/**
* $ go run src/test/cycle.go < data/tinyCG.txt
* 1 0 2 1
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	finder := NewCycle(g)

	if finder.HasCycle() {
		c := finder.Cycle()
		for w := c.Next(); w != nil; w = c.Next() {
			fmt.Print(w, " ")
		}
		fmt.Println()
		//fmt.Println(c)
	} else {
		fmt.Println("Graph is acyclic")
	}

}
