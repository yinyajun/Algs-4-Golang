package main

import (
	"os"

	"algs4/graph"
	"strconv"
	"math/rand"
	"fmt"
)

/**
* $ go run src/test/edgeWeightedDirectedCycle.go 5 6 2
* extra: 3 0
* extra: 3 3
* 5 8
* 0:
* 1: 1 -> 0  0.30 1 -> 0  0.36 1 -> 0  0.68 1 -> 0  0.52
* 2: 2 -> 1  0.98
* 3: 3 -> 3  0.53 3 -> 0  0.54
* 4: 4 -> 3  0.47
*
* Cycle: 3 -> 3  0.53
*
*
* $ go run src/test/edgeWeightedDirectedCycle.go 5 6 1
* extra: 3 0
* 5 7
* 0:
* 1: 1 -> 0  0.30 1 -> 0  0.36 1 -> 0  0.68 1 -> 0  0.52
* 2: 2 -> 1  0.98
* 3: 3 -> 0  0.54
* 4: 4 -> 3  0.47
*
* No directed cycle
*
*
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	// create random DAG with V vertices and E edges; then add F random edges
	V, _ := strconv.Atoi(os.Args[1])
	E, _ := strconv.Atoi(os.Args[2])
	F, _ := strconv.Atoi(os.Args[3])
	G := graph.NewEdgeWeightedDigraph(V)
	vertices := make([]int, V)
	for i := 0; i < V; i++ {
		vertices[i] = i
	}
	rand.Shuffle(len(vertices), func(i, j int) { vertices[i], vertices[j] = vertices[j], vertices[i] })
	for i := 0; i < E; i++ {
		var v, w int
		for {
			v = rand.Intn(V)
			w = rand.Intn(V)
			if v > w {
				break
			}
		}
		weight := rand.Float64()
		G.AddEdge(graph.NewDirectedEdge(v, w, weight))
	}

	// add F extra edges
	for i := 0; i < F; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := rand.Float64()
		fmt.Println("extra:", v, w)
		G.AddEdge(graph.NewDirectedEdge(v, w, weight))
	}
	fmt.Println(G)

	// find a directed cycle
	finder := graph.NewEdgeWeightedDirectedCycle(G)
	if finder.HasCycle() {
		fmt.Print("Cycle: ")
		edges := finder.Cycle()
		for e := edges.Next(); e != nil; e = edges.Next() {
			fmt.Print(e, " ")
		}
		fmt.Println()
	} else { // or give topologial sort
		fmt.Println("No directed cycle")
	}

}
