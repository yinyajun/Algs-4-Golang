package graph

import . "algs4/bag"

// undirected graph
type Graph interface {
	V() int           // vertex num
	E() int           // edge num
	AddEdge(v, w int) // add edge v-w
	Adj(v int) Bag    // all adjacent vertex of v
}

type Search interface {
	Search(g graph, s int) // find all vertices connected to s
	Marked(v int) bool     // is v and s connected?
	Count() int            // number of vertices connected to s
}
