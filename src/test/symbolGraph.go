package main

import (
	"algs4/graph"
	"bufio"
	"fmt"
	"os"
	. "util"
)

/**
*	go run src/test/symbolGraph.go "data/routes.txt" " "
*	JFK
*		ORD
*		ATL
*		MCO
*	LAX
*		LAS
*		PHX
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := graph.NewSymbolGraph(filename, delimiter)
	g := sg.Graph()
	in := NewInWithSplitFunc(os.Stdin, bufio.ScanLines)

	for in.HasNext() {
		source := in.ReadLine()
		if sg.Contains(source) {
			s := sg.Index(source)
			gen := g.Adj(s)
			for v := gen.Next(); v != nil; v = gen.Next() {
				fmt.Println(" ", sg.Name(v.(int)))
			}
		} else {
			fmt.Println("input not contains '", source, "'")
		}
	}
}
