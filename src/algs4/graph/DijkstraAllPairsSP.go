package graph

import "util"

/**
* represents a data type for solving the all-pairs shortest paths problem in edge-weighted digraphs
* where the edge weights are non-negative.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DijkstraAllPairsSP struct {
	all []*DijkstraSP
}

func NewDijkstraAllPairsSP(g *EdgeWeightedDigraph) *DijkstraAllPairsSP {
	m := &DijkstraAllPairsSP{}
	m.all = make([]*DijkstraSP, g.V())
	for v := 0; v < g.V(); v++ {
		m.all[v] = NewDijkstraSP(g, v)
	}
	return m
}

func (m *DijkstraAllPairsSP) Path(s, t int) util.Iterator {
	m.validateVertex(s)
	m.validateVertex(t)
	return m.all[s].PathTo(t)
}

func (m *DijkstraAllPairsSP) HasPath(s, t int) bool {
	m.validateVertex(s)
	m.validateVertex(t)
	return m.Dist(s, t) < POSTIVE_INFINITY
}

func (m *DijkstraAllPairsSP) Dist(s, t int) float64 {
	m.validateVertex(s)
	m.validateVertex(t)
	return m.all[s].DistTo(t)
}

func (m *DijkstraAllPairsSP) validateVertex(v int) {
	V := len(m.all)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
