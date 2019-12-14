package graph

import (
	. "util"
)

// undirected graph
type Graph interface {
	V() int             // vertex num
	E() int             // edge num
	AddEdge(v, w int)   // add edge v-w
	Adj(v int) Iterator // all adjacent vertex of v
}

type Search interface {
	Marked(v int) bool // is v and s connected?
	Count() int        // number of vertices connected to s
}

type Paths interface {
	HasPathTo(v int) bool
	PathTo(v int) Iterator // path from s to v, empty if not exists
}

type CC interface {
	Connected(v, w int) bool
	Count() int
	Id(v int) int
}

type SymbolGraph interface {
	Contains(key string) bool
	Index(key string) int
	Name(v int) string
	Graph() *graph
}

type Digraph interface {
	V() int             // vertex num
	E() int             // edge num
	AddEdge(v, w int)   // add edge v-w
	Adj(v int) Iterator // all adjacent vertex of v
	Reverse() *digraph
}

type HasCycle interface {
	HasCycle() bool
	Cycle() Iterator
}

type Topological interface {
	IsDAG() bool
	Order() Iterator
}
