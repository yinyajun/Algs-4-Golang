package graph

import (
	. "algs4/priorityQueue"
	"algs4/stack"
	"fmt"
	"math"
	"util"
)

/**
* a data type for solving the single-source shortest paths problem in edge-weighted digraphs
* where the edge weights are nonnegative.
*
* This implementation  uses Dijkstra's algorithm with a binary heap.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DijkstraSP struct {
	distTo []float64       // distTo[v] = distance  of shortest s->v path
	edgeTo []*DirectedEdge // edgeTo[v] = last edge on shortest s->v path
	pq     *IndexMinPQ     // priority queue of vertices
}

func NewDijkstraSP(g *EdgeWeightedDigraph, s int) *DijkstraSP {
	m := &DijkstraSP{}
	edges := g.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		if e.(*DirectedEdge).Weight() < 0 {
			panic("NewDijkstraSP: edge has negative weight")
		}
	}

	m.distTo = make([]float64, g.V())
	m.edgeTo = make([]*DirectedEdge, g.V())
	m.validateVertex(s)
	for v := 0; v < g.V(); v++ {
		m.distTo[v] = POSTIVE_INFINITY
	}

	m.pq = NewIndexMinPQ(g.V())

	m.distTo[s] = 0.0
	// relax vertices in order of distance from s
	m.pq.Insert(s, m.distTo[s])
	for !m.pq.IsEmpty() {
		v := m.pq.DelMin()
		vAdj := g.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			m.relax(e.(*DirectedEdge))
		}
	}

	// check optimality conditions
	if !m.check(g, s) {
		panic("check failed")
	}
	return m
}

func (m *DijkstraSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if m.distTo[v]+e.Weight() < m.distTo[w] {
		m.distTo[w] = m.distTo[v] + e.Weight()
		m.edgeTo[w] = e
		if m.pq.Contains(w) {
			m.pq.DecreaseKey(w, m.distTo[w])
		} else {
			m.pq.Insert(w, m.distTo[w])
		}
	}
}

func (m *DijkstraSP) DistTo(v int) float64 {
	m.validateVertex(v)
	return m.distTo[v]
}

func (m *DijkstraSP) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.distTo[v] < POSTIVE_INFINITY
}

func (m *DijkstraSP) PathTo(v int) util.Iterator {
	m.validateVertex(v)
	if !m.HasPathTo(v) {
		return nil
	}
	path := stack.NewStack()
	for e := m.edgeTo[v]; e != nil; e = m.edgeTo[e.From()] {
		path.Push(e)
	}
	return path.Iterate()
}

func (m *DijkstraSP) check(g *EdgeWeightedDigraph, s int) bool {
	// check that edge weights are non-negative
	edges := g.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		if e.(*DirectedEdge).Weight() < 0 {
			fmt.Println("negative edge weight detected")
			return false
		}
	}
	// check that distTo[v] and edgeTo[v] are consistent
	if m.distTo[s] != 0.0 || m.edgeTo[s] != nil {
		fmt.Println("distTo[s] and edgeTo[s] inconsistent")
		return false
	}
	for v := 0; v < g.V(); v++ {
		if v == s {
			continue
		}
		if m.edgeTo[v] == nil && math.Abs(m.distTo[v]-POSTIVE_INFINITY) > 1E-10 {
			fmt.Println("distTo[] and edgeTo[] inconsistent")
			return false
		}
	}
	// check that all edges e = v->w satisfy distTo[w] <= distTo[v] + e.weight()
	for v := 0; v < g.V(); v++ {
		vAdj := g.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			w := e.(*DirectedEdge).To()
			if m.distTo[v]+e.(*DirectedEdge).Weight() < m.distTo[w] {
				fmt.Printf("edge %s not relaxed", e.(*DirectedEdge))
				return false
			}
		}
	}
	// check that all edges e = v->w on SPT satisfy distTo[w] == distTo[v] + e.weight()
	for w := 0; w < g.V(); w++ {
		if m.edgeTo[w] == nil {
			continue
		}
		e := m.edgeTo[w]
		v := e.From()
		if w != e.To() {
			return false
		}
		if math.Abs(m.distTo[v]+e.Weight()-m.distTo[w]) > 1E-10 {
			fmt.Printf("edge %s on shortest path not tight.", e)
			return false
		}
	}
	return true
}

func (m *DijkstraSP) validateVertex(v int) {
	V := len(m.distTo)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
