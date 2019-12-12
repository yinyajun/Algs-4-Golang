package main

import (
	"algs4/graph"
	"fmt"
	"os"
	. "util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := graph.NewSymbolGraph(filename, delimiter)
	graph := sg.Graph()
	in := NewIn(os.Stdin)

	for in.HasNext() {
		source := in.ReadString()
		if sg.Contains(source) {
			s := sg.Index(source)
			gen := graph.Adj(s).Yield()
			for hasNext, v := gen(); hasNext; hasNext, v = gen() {
				fmt.Println(" ", sg.Name(v.(int)))
			}
		} else {
			fmt.Println("input not contains '", source, "'")
		}
	}
}
