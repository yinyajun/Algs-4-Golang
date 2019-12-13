package main

import (
	"algs4/graph"
	"fmt"
	"os"
	. "util"
	"bufio"
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
	g := sg.Graph()
	in := NewInWithSplitFunc(os.Stdin, bufio.ScanLines)

	for in.HasNext() {
		source := in.ReadLine()
		if sg.Contains(source) {
			s := sg.Index(source)
			gen := g.Adj(s)
			for hasNext, v := gen(); hasNext; hasNext, v = gen() {
				fmt.Println(" ", sg.Name(v.(int)))
			}
		} else {
			fmt.Println("input not contains '", source, "'")
		}
	}
}