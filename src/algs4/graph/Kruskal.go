package graph

import (
	. "algs4/priorityQueue"
	. "algs4/queue"
	. "algs4/unionFind"
	"fmt"
	"math"
	. "util"
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

	m.mst = NewQueue()

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

func (m *KruskalMST) Edges() Iterator { return m.mst.Iterate() }

func (m *KruskalMST) Weight() float64 { return m.weight }

func (m *KruskalMST) check(g *EdgeWeightedGraph) bool {
	// check weight
	var totalWeight float64
	edges := m.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		totalWeight += e.(*Edge).Weight()
	}
	if math.Abs(totalWeight-m.Weight()) > 1E-10 {
		fmt.Println("Weight of edges does not equal Weight()", totalWeight, m.Weight())
		return false
	}

	// check that it is acyclic (use union find to determine cycle existence)
	uf := NewPathComWeightedQU(g.V())
	edges = m.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		v := e.(*Edge).Either()
		w := e.(*Edge).Other(v)
		if uf.Connected(v, w) {
			fmt.Println("Not a forest", v, w)
			return false
		}
		uf.Union(v, w)
	}

	// check that it is a spanning forest
	allEdges := g.Edges()
	for e := allEdges.Next(); e != nil; e = allEdges.Next() {
		v := e.(*Edge).Either()
		w := e.(*Edge).Other(v)
		if !uf.Connected(v, w) {
			fmt.Println("Not a spanning forest")
			return false
		}
	}

	// check that it is a minimal spanning forest (cut optimality conditions)
	edges = m.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {

		// all edges in MST except e
		uf := NewPathComWeightedQU(g.V())
		es := m.Edges()
		for f := es.Next(); f != nil; f = es.Next() {
			x := f.(*Edge).Either()
			y := f.(*Edge).Other(x)
			if f != e {
				uf.Union(x, y)
			}
		}

		// check that e is min weight edge in crossing cut
		gEdges := g.Edges()
		for f := gEdges.Next(); f != nil; f = gEdges.Next() {
			x := f.(*Edge).Either()
			y := f.(*Edge).Other(x)
			if !uf.Connected(x, y) { // this edge must be crossing cut
				if f.(*Edge).Weight() < e.(*Edge).Weight() {
					fmt.Println("Edge", f, "violates cut optimality conditions")
					return false
				}
			}
		}
	}
	return true
}
