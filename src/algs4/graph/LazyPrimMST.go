package graph

import (
	. "algs4/queue"
	. "algs4/priorityQueue"
)

/**
* computing a minimum spanning tree in an Edge-weighted Graph.
*
* This implementation uses a lazy version of Prim's algorithm
* with a binary heap of edges.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type LazyPrimMST struct {
	marked []bool // marked[v] = true iff v on tree
	mst    *Queue // edges in the MST
	pq     *MinPQ // edges with one endpoint in tree
	weight int    // total weight of MST
}

func NewLazyPrimMST(g *EdgeWeightedGraph) *LazyPrimMST {
	m := &LazyPrimMST{}
	m.marked = make([]bool, g.V())
	m.mst = NewQueue()
	m.pq = NewMinPQ()

	return m
}
