package graph

import (
	. "algs4/bag"
	. "util"
)

/**
*
* Digraph
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type digraph struct {
	v        int    // number of vertices in this digraph
	e        int    // number of edges in this digraph
	adj      []*Bag // adj[v] = adjacency list for vertex v
	indegree []int  // indegree[v] = indegree of vertex v
}

func NewDigraph(V int) *digraph {
	if V < 0 {
		panic("NewDigraph: invalid V")
	}
	dg := &digraph{}
	dg.adj = make([]*Bag, V)
	for idx := range dg.adj {
		dg.adj[idx] = &Bag{}
	}
	dg.indegree = make([]int, V)
	return dg
}

func NewDigraphWithIn(in In) *digraph {
	v := in.ReadInt()
	dg := NewDigraph(v)
	e := in.ReadInt()
	if e < 0 {
		panic("NewDigraphWithIn: invalid e")
	}
	for i := 0; i < e; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		dg.AddEdge(v, w)
	}
	return dg
}
