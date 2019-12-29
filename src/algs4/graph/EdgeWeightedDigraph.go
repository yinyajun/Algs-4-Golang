package graph

import (
	. "algs4/bag"
	"algs4/stack"
	"fmt"
	"strings"
)
import . "util"

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type EdgeWeightedDigraph struct {
	v        int    // number of vertices in this digraph
	e        int    // number of edges in this digraph
	adj      []*Bag // adj[v] = adjacency list for vertex v
	indegree []int  // indegree[v] = indegree of vertex v
}

func NewEdgeWeightedDigraph(v int) *EdgeWeightedDigraph {
	m := &EdgeWeightedDigraph{}
	m.v = v
	m.adj = make([]*Bag, v)
	for idx := range m.adj {
		m.adj[idx] = NewBag()
	}
	m.indegree = make([]int, v)
	return m
}

func NewEdgeWeightedDigraphWithIn(in *In) *EdgeWeightedDigraph {
	m := NewEdgeWeightedDigraph(in.ReadInt())
	e := in.ReadInt()
	if e < 0 {
		panic("Number of edges must be non-negative.")
	}
	for i := 0; i < e; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		m.validateVertex(v)
		m.validateVertex(w)
		weight := in.ReadFloat()
		m.AddEdge(NewDirectedEdge(v, w, weight))
	}
	return m
}

func NewEdgeWeightedDigraphWithGraph(g *EdgeWeightedDigraph) *EdgeWeightedDigraph {
	m := NewEdgeWeightedDigraph(g.V())
	m.e = g.E()
	for v := 0; v < g.V(); v++ {
		m.indegree[v] = g.Indegree(v)
	}
	for v := 0; v < g.V(); v++ {
		reverse := stack.NewStack()
		vAdj := g.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			reverse.Push(e)
		}
		it := reverse.Iterate()
		for e := it.Next(); e != nil; e = it.Next() {
			m.adj[v].Add(e)
		}
	}
	return m
}

func (m *EdgeWeightedDigraph) V() int { return m.v }

func (m *EdgeWeightedDigraph) E() int { return m.e }

func (m *EdgeWeightedDigraph) validateVertex(v int) {
	if v < 0 || v >= m.v {
		panic("validateVertex: invalid vertex")
	}
}

func (m *EdgeWeightedDigraph) AddEdge(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	m.validateVertex(v)
	m.validateVertex(w)
	m.adj[v].Add(e)
	m.indegree[v]++
	m.e++
}

func (m *EdgeWeightedDigraph) Adj(v int) Iterator {
	m.validateVertex(v)
	return m.adj[v].Iterate()
}

func (m *EdgeWeightedDigraph) Outdegree(v int) int {
	m.validateVertex(v)
	return m.adj[v].Size()
}

func (m *EdgeWeightedDigraph) Indegree(v int) int {
	m.validateVertex(v)
	return m.indegree[v]
}

func (m *EdgeWeightedDigraph) Edges() Iterator {
	list := NewBag()
	for v := 0; v < m.v; v++ {
		vAdj := m.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			list.Add(e)
		}
	}
	return list.Iterate()
}

func (m *EdgeWeightedDigraph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d %d\n", m.v, m.e))
	for v := 0; v < m.v; v++ {
		s.WriteString(fmt.Sprintf("%d: ", v))
		vAdj := m.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			s.WriteString(fmt.Sprintf("%v ", e.(*DirectedEdge)))
		}
		s.WriteString("\n")
	}
	return s.String()

}
