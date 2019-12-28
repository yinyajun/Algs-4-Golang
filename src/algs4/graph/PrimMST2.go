package graph

import (
	. "algs4/priorityQueue"
	"algs4/queue"
	"algs4/unionFind"
	"fmt"
	"math"
	"util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type PrimMST2 struct {
	edgeTo []*Edge
	marked []bool
	pq     *IndexMinPQ
}

func NewPrimMST2(g *EdgeWeightedGraph) *PrimMST2 {
	m := &PrimMST2{}
	m.edgeTo = make([]*Edge, g.V())
	m.marked = make([]bool, g.V())
	m.pq = NewIndexMinPQ(g.V())

	for v := 0; v < g.V(); v++ {
		if !m.marked[v] {
			m.prim(g, v)
		}
	}
	//check optimality conditions
	if !m.check(g) {
		panic("NewPrimMST: check failed")
	}
	return m
}

func (m *PrimMST2) prim(g *EdgeWeightedGraph, s int) {
	m.pq.Insert(s, 0)
	for !m.pq.IsEmpty() {
		v := m.pq.DelMin()
		m.scan(g, v)
	}
}

func (m *PrimMST2) scan(g *EdgeWeightedGraph, v int) {
	m.marked[v] = true
	vAdj := g.Adj(v)
	for e := vAdj.Next(); e != nil; e = vAdj.Next() {
		w := e.(*Edge).Other(v)
		if m.marked[w] {
			continue // v-w is obsolete edge
		}
		if m.edgeTo[w] == nil || (m.edgeTo[w] != nil && e.(*Edge).weight < m.edgeTo[w].Weight()) {
			m.edgeTo[w] = e.(*Edge)
			if m.pq.Contains(w) {
				m.pq.DecreaseKey(w, m.edgeTo[w].Weight())
			} else {
				m.pq.Insert(w, m.edgeTo[w].Weight())
			}
		}
	}

}

func (m *PrimMST2) Edges() util.Iterator {
	mst := queue.NewQueue()
	for _, e := range m.edgeTo {
		if e != nil {
			mst.Enqueue(e)
		}
	}
	return mst.Iterate()
}

func (m *PrimMST2) Weight() float64 {
	var weight float64
	for _, e := range m.edgeTo {
		if e == nil {
			continue
		}
		weight += e.Weight()
	}
	return weight
}

func (m *PrimMST2) check(g *EdgeWeightedGraph) bool {

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
