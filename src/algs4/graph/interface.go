package main

// undirected graph
type Graph interface {
	V() int           // vertex num
	E() int           // edge num
	addEdge(v, w int) // add edge v-w
	adj(v int) []int  // all adjacent vertex of v
}
