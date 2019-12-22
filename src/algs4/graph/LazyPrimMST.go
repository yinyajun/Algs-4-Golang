package graph

import (
	. "algs4/queue"
	. "algs4/priorityQueue"
	"util"
	"math"
	"fmt"
	"algs4/unionFind"
)

/**
* computing a minimum spanning tree in an Edge-weighted Graph.
*
* This implementation uses a lazy version of Prim's algorithm
* with a binary heap of edges.
*
* 优先队列保存所有横切边， 顶点索引数组标记mst顶点，队列来保存mst
* 延时实现会在优先队列中保留失效边
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type LazyPrimMST struct {
	marked []bool  // marked[v] = true iff v on tree
	mst    *Queue  // edges in the MST
	pq     *MinPQ  // edges with one endpoint in tree
	weight float64 // total weight of MST
}

func NewLazyPrimMST(g *EdgeWeightedGraph) *LazyPrimMST {
	m := &LazyPrimMST{}
	m.marked = make([]bool, g.V())
	m.mst = NewQueue()
	m.pq = NewMinPQwithCapAndCom(1, EdgeComparator{})

	for v := 0; v < g.V(); v++ { // run Prim from all vertices to
		if !m.marked[v] { // get a minimum spanning forest
			m.prim(g, v)
		}
	}

	if !m.check(g) {
		panic("LazyPrimMST failed")
	}
	return m
}

func (m *LazyPrimMST) prim(g *EdgeWeightedGraph, s int) {
	m.scan(g, s)
	for !m.pq.IsEmpty() {
		e := m.pq.DelMin().(*Edge) // smallest edge on pq
		v := e.Either()            // two endpoints
		w := e.Other(v)
		if !m.marked[v] && !m.marked[w] {
			panic("prim: neither endpoint is part of tree")
		}
		if m.marked[v] && m.marked[w] { // lazy, both v and w already scanned
			continue
		}
		m.mst.Enqueue(e) // add e to MST
		m.weight += e.Weight()

		if m.mst.Size()+1 == g.V() { // early stop
			break
		}

		if !m.marked[v] { // v becomes part of tree
			m.scan(g, v)
		}
		if !m.marked[w] { // w becomes part of tree
			m.scan(g, w)
		}
	}
}

// add all edges e incident to v onto pq if the other endpoint has not yet been scanned
func (m *LazyPrimMST) scan(g *EdgeWeightedGraph, v int) {
	if m.marked[v] {
		panic("scan: v has been visited")
	}
	m.marked[v] = true

	vAdj := g.Adj(v)
	for e := vAdj.Next(); e != nil; e = vAdj.Next() {
		if !m.marked[e.(*Edge).Other(v)] {
			m.pq.Insert(e)
		}
	}
}

func (m *LazyPrimMST) Edges() util.Iterator { return m.mst.Iterate() }

func (m *LazyPrimMST) Weight() float64 { return m.weight }

// check optimality conditions (takes time proportional to E V lg* V)
func (m *LazyPrimMST) check(g *EdgeWeightedGraph) bool {
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
	uf := unionFind.NewPathComWeightedQU(g.V())
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
		uf := unionFind.NewPathComWeightedQU(g.V())
		mstEdges := m.mst.Iterate()
		for f := mstEdges.Next(); f != nil; f = mstEdges.Next() {
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
