package graph

import . "algs4/bag"

// undirected graph
type Graph interface {
	V() int           // vertex num
	E() int           // edge num
	AddEdge(v, w int) // add edge v-w
	Adj(v int) Bag    // all adjacent vertex of v
}
