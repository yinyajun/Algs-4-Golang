package main

import (
	"algs4/graph"
	"bufio"
	"fmt"
	"os"

	. "util"
)

/**
* $ go run src/test/symbolDigraph.go "data/routes.txt" " "
* JFK
*   ORD
*   ATL
*   MCO
* ATL
*   MCO
*   HOU
* LAX
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sd := graph.NewSymbolDigraph(filename, delimiter)
	g := sd.Digraph()

	in := NewInWithSplitFunc(os.Stdin, bufio.ScanLines)
	for in.HasNext() {
		source := in.ReadLine()
		if sd.Contains(source) {
			s := sd.Index(source)
			gen := g.Adj(s)
			for hasNext, v := gen(); hasNext; hasNext, v = gen() {
				fmt.Println(" ", sd.Name(v.(int)))
			}
		} else {
			fmt.Println("input not contains '", source, "'")
		}
	}
}
