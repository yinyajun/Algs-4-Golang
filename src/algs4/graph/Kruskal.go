package graph

import (
	. "algs4/priorityQueue"
	. "algs4/queue"
	. "algs4/unionFind"
)

/**
* a data type for computing a minimum spanning tree in an edge-weighted graph.
*
* This implementation uses Krusal's algorithm and the union-find data type.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type KruskalMST struct {
	mst    *Queue  // edges in MST
	weight float64 // weight of MST
	pq     *MinPQ
}

func NewKruskalMST(g *EdgeWeightedGraph) *KruskalMST {
	m := &KruskalMST{}

	// heapify an array, it is more efficient
	m.pq = NewMinPQwithArray(g.EdgesArray())

	// run greedy algorithm
	uf := NewPathComWeightedQU(g.V())
	for !m.pq.IsEmpty() && m.mst.Size() < (g.V()-1) {
		e := m.pq.DelMin().(*Edge)
		v := e.Either()
		w := e.Other(v)
		if !uf.Connected(v, w) { // v-w does not create a cycle
			uf.Union(v, w)   // merge v and w components
			m.mst.Enqueue(e) // add edge e to mst
			m.weight += e.Weight()
		}
	}

	// check optimality conditions
	if !m.check(g) {
		panic("check failed")
	}

	return m
}
