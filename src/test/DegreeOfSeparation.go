package main

import (
	"algs4/graph"
	"fmt"
	"os"
	"util"
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
	source := os.Args[3]

	sg := graph.NewSymbolGraph(filename, delimiter)
	g := sg.Graph()
	if !sg.Contains(source) {
		fmt.Println(source, "not in database.")
		return
	}

	s := sg.Index(source)
	bfs := graph.NewBreadthFirstPaths(g, s)

	in := util.NewIn(os.Stdin)
	sink := in.ReadString()
	if sg.Contains(sink) {
		t := sg.Index(sink)
		if bfs.HasPathTo(t) {
			gen := bfs.PathTo(t)
			for hasNext, v := gen(); hasNext; hasNext, v = gen() {
				fmt.Println(" ", sg.Name(v.(int)))
			}
		} else {
			fmt.Println("Not connected")
		}
	} else {
		fmt.Println(" Not in databases.")
	}
}