package graph

import (
	"algs4/stack"
	"util"
)

/**
*  represents a data type for solving the single-source shortest paths problem in edge-weighted
*  directed acyclic graphs (DAGs). The edge weights can be positive, negative, or zero.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type AcyclicSP struct {
	distTo []float64
	edgeTo []*DirectedEdge
}

func NewAcyclicSP(g *EdgeWeightedDigraph, s int) *AcyclicSP {
	m := &AcyclicSP{}
	m.distTo = make([]float64, g.V())
	m.edgeTo = make([]*DirectedEdge, g.V())

	m.validateVertex(s)

	for v := 0; v < g.V(); v++ {
		m.distTo[v] = POSTIVE_INFINITY
	}
	m.distTo[s] = 0.0

	// visit vertices in topological order
	order := NewTopologicalEWD(g)
	if !order.HasOrder() {
		panic("Digraph is not acyclic")
	}
	orders := order.Order()
	for v := orders.Next(); v != nil; v = orders.Next() {
		vAdjEdges := g.Adj(v.(int))
		for e := vAdjEdges.Next(); e != nil; e = vAdjEdges.Next() {
			m.relax(e.(*DirectedEdge))
		}
	}
	return m
}

func (m *AcyclicSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if m.distTo[w] > m.distTo[v]+e.Weight() {
		m.distTo[w] = m.distTo[v] + e.Weight()
		m.edgeTo[w] = e
	}
}

func (m *AcyclicSP) DistTo(v int) float64 {
	m.validateVertex(v)
	return m.distTo[v]
}

func (m *AcyclicSP) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.distTo[v] < POSTIVE_INFINITY
}

func (m *AcyclicSP) PathTo(v int) util.Iterator {
	m.validateVertex(v)
	path := stack.NewStack()
	for e := m.edgeTo[v]; e != nil; e = m.edgeTo[e.From()] {
		path.Push(e)
	}
	return path.Iterate()

}

func (m *AcyclicSP) validateVertex(v int) {
	V := len(m.distTo)
	if v < 0 || v >= V {
		panic("invalid vertex")
	}
}
