package main

import (
	"algs4/graph"
	"fmt"
	"os"
)

/**
*$ go run src/test/topological.go "data/jobs.txt" "/"
* Calculus
* Linear Algebra
* Introduction to CS
* Advanced Programming
* Algorithms
* Theoretical CS
* Artificial Intelligence
* Robotics
* Machine Learning
* Neural Networks
* Databases
* Scientific Computing
* Computational Biology
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := graph.NewSymbolDigraph(filename, delimiter)

	topological := graph.NewTopological(sg.Digraph())
	order := topological.Order()
	fmt.Println(order)

	for ok, v := order(); ok; ok, v = order() {
		fmt.Println(sg.Name(v.(int)))
	}
}
