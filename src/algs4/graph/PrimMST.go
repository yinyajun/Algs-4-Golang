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
* computing a minimum spanning tree in an Edge-weighted Graph.
*
* This implementation uses Prim's algorithm with a indexed binary heap.
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

const POSTIVE_INFINITY = float64(^uint(0) >> 1)

type PrimMST struct {
	edgeTo []*Edge   // edgeTo[v] = shortest edge from tree vertex to non-tree vertex
	distTo []float64 // distTo[v] = weight of shortest such edge
	marked []bool    // marked[v] = true if v on tree, false otherwise
	pq     *IndexMinPQ
}

func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	m := &PrimMST{}
	m.edgeTo = make([]*Edge, g.V())
	m.distTo = make([]float64, g.V())
	for idx := range m.distTo {
		m.distTo[idx] = POSTIVE_INFINITY
	}
	m.marked = make([]bool, g.V())
	m.pq = NewIndexMinPQ(g.V())

	for v := 0; v < g.V(); v++ { // run Prim from all vertices to
		if !m.marked[v] { // get a minimum spanning forest
			m.prim(g, v)
		}
	}

	//check optimality conditions
	if !m.check(g) {
		panic("NewPrimMST: check failed")
	}
	return m
}

// run Prim's algorithm in graph G, starting from vertex s
func (m *PrimMST) prim(g *EdgeWeightedGraph, s int) {
	m.distTo[s] = 0
	m.pq.Insert(s, m.distTo[s])
	for !m.pq.IsEmpty() {
		v := m.pq.DelMin()
		m.scan(g, v)
	}
}

func (m *PrimMST) scan(g *EdgeWeightedGraph, v int) {
	m.marked[v] = true
	vAdj := g.Adj(v)
	for e := vAdj.Next(); e != nil; e = vAdj.Next() {
		w := e.(*Edge).Other(v)
		if m.marked[w] {
			continue // v-w is obsolete edge
		}
		if e.(*Edge).Weight() < m.distTo[w] {
			m.distTo[w] = e.(*Edge).Weight()
			m.edgeTo[w] = e.(*Edge)
			if m.pq.Contains(w) {
				m.pq.DecreaseKey(w, m.distTo[w])
			} else {
				m.pq.Insert(w, m.distTo[w])
			}
		}
	}
}

func (m *PrimMST) Edges() util.Iterator {
	mst := queue.NewQueue()
	for _, e := range m.edgeTo {
		if e != nil {
			mst.Enqueue(e)
		}
	}
	return mst.Iterate()
}

func (m *PrimMST) Weight() float64 {
	var weight float64
	for _, e := range m.edgeTo {
		if e == nil {
			continue
		}
		weight += e.Weight()
	}
	return weight
}

func (m *PrimMST) check(g *EdgeWeightedGraph) bool {

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
