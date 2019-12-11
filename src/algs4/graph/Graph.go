package main

import (
	"fmt"

	. "algs4/bag"
)

type graph struct {
	V   int   // # vertex
	E   int   // # edge
	adj []Bag //Adjacency list
}

func (g *graph) newGraph(V int) {
	if V < 0 {
		panic("newGraph: V should be non-negative")
	}
	g.V = V
	g.adj = make([]Bag, V)
}

//func (g *graph) newGraphWithIn(V int) {}

func main() {
	c := make([]Bag, 5)
	fmt.Println(c)
}
