package graph

import . "util"

// undirected Graph
type graph interface {
	V() int             // vertex num
	E() int             // Edge num
	AddEdge(v, w int)   // add Edge v-w
	Adj(v int) Iterator // all adjacent vertex of v
}

type search interface {
	Marked(v int) bool // is v and s connected?
	Count() int        // number of vertices connected to s
}

type paths interface {
	HasPathTo(v int) bool
	PathTo(v int) Iterator // path from s to v, empty if not exists
}

type cc interface {
	Connected(v, w int) bool
	Count() int
	Id(v int) int
}

type symbolGraph interface {
	Contains(key string) bool
	Index(key string) int
	Name(v int) string
	Graph() *Graph
}

type digraph interface {
	V() int             // vertex num
	E() int             // Edge num
	AddEdge(v, w int)   // add Edge v-w
	Adj(v int) Iterator // all adjacent vertex of v
	Reverse() *Digraph
}

type hasCycle interface {
	HasCycle() bool
	Cycle() Iterator
}

type topological interface {
	IsDAG() bool
	Order() Iterator
}

type transitiveClosure interface {
	Reachable(v, w int) bool
}

type edge interface {
	Weight() float64
	Either() int
	Other(v int) int
	CompareTo(that *Edge) bool
}

type edgeWeightedGraph interface {
	V() int
	E() int
	AddEdge(e *Edge)
	Adj(v int) Iterator
	Edges() Iterator
}

type mst interface {
	Edges() Iterator
	Weight() float64
}
