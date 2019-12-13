package main

import (
	"fmt"
	"os"
	. "algs4/graph"
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
		gen := finder.Cycle()
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			fmt.Print(w, " ")
		}
		fmt.Println()
		//fmt.Println(gen)
	} else {
		fmt.Println("Graph is acyclic")
	}

}
