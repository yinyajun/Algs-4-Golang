package main

import (
	"algs4/graph"
	"bufio"
	"fmt"
	"os"
	"util"
)

/**
* $ go run src/test/DegreeOfSeparation.go "data/movies.txt" "/" "Bacon, Kevin"
* Kidman, Nicole
*   Bacon, Kevin
*   Woodsman, The (2004)
*   Grier, David Alan
*   Bewitched (2005)
*   Kidman, Nicole
* Grant, Cary
*   Bacon, Kevin
*   Planes, Trains & Automobiles (1987)
*   Martin, Steve (I)
*   Dead Men Don't Wear Plaid (1982)
*   Grant, Cary
*
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

	in := util.NewInWithSplitFunc(os.Stdin, bufio.ScanLines)
	for in.HasNext() {
		sink := in.ReadLine()
		if sg.Contains(sink) {
			t := sg.Index(sink)
			if bfs.HasPathTo(t) {
				path := bfs.PathTo(t)
				for v := path.Next(); v != nil; v = path.Next() {
					fmt.Println(" ", sg.Name(v.(int)))
				}
			} else {
				fmt.Println("Not connected")
			}
		} else {
			fmt.Println("\tNot in databases.")
		}
	}

}
