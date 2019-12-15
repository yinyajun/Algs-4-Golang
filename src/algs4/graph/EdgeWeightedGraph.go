package graph

import (
	"algs4/bag"
	"math/rand"
	"util"
	"algs4/stack"
	"strings"
	"fmt"
)

/**
*
* represents an Edge-weighted Graph of vertices named 0 through V – 1, where each
* undirected Edge is of type Edge and has a real-valued weight.
*
* This implementation uses an adjacency-lists representation
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type EdgeWeightedGraph struct {
	v   int
	e   int
	adj []*bag.Bag
}

// Initializes an empty Edge-weighted Graph with v vertices and 0 edges.
func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	m := &EdgeWeightedGraph{}
	if v < 0 {
		panic("NewEdgeWeightedGraph: invalid v")
	}
	m.v = v
	m.adj = make([]*bag.Bag, v)
	for idx := range m.adj {
		m.adj[idx] = bag.NewBag()
	}
	return m
}

// Initializes a random Edge-weighted Graph with V vertices and E edges.
func NewEdgeWeightedGraphRandomly(V, E int) *EdgeWeightedGraph {
	m := NewEdgeWeightedGraph(V)
	if E < 0 {
		panic("NewEdgeWeightedGraphRandomly: invalid e")
	}
	for i := 0; i < E; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := rand.Float64()
		edge := NewEdge(v, w, weight)
		m.AddEdge(edge)
	}
	return m
}

// Initializes an Edge-weighted Graph from an input stream.
func NewEdgeWeightedGraphWithIn(in *util.In) *EdgeWeightedGraph {
	v := in.ReadInt()
	m := NewEdgeWeightedGraph(v)
	e := in.ReadInt()
	if e < 0 {
		panic("NewEdgeWeightedGraphWithIn: invalid e")
	}
	for i := 0; i < e; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		weight := in.ReadFloat()
		e := NewEdge(v, w, weight)
		m.AddEdge(e)
	}
	return m
}

// Initializes a new Edge-weighted Graph that is a deep copy of g
func NewEdgeWeightedGraphwithG(g *EdgeWeightedGraph) *EdgeWeightedGraph {
	m := NewEdgeWeightedGraph(g.V())
	m.e = g.E()
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

func (m *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	m.validateVertex(v)
	m.validateVertex(w)
	m.adj[v].Add(e)
	m.adj[w].Add(e)
	m.e++
}

func (m *EdgeWeightedGraph) validateVertex(v int) {
	V := len(m.adj)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *EdgeWeightedGraph) V() int { return m.v }

func (m *EdgeWeightedGraph) E() int { return m.e }

func (m *EdgeWeightedGraph) Adj(v int) util.Iterator {
	m.validateVertex(v)
	return m.adj[v].Iterate()
}

func (m *EdgeWeightedGraph) Degree(v int) int {
	m.validateVertex(v)
	return m.adj[v].Size()
}

func (m *EdgeWeightedGraph) Edges() util.Iterator {
	list := bag.NewBag()
	for v := 0; v < m.v; v++ {
		selfLoops := 0
		vAdj := m.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			if e.(*Edge).Other(v) > v {
				list.Add(e)
			} else if e.(*Edge).Other(v) == v {
				if selfLoops%2 == 0 {
					list.Add(e)
				}
				selfLoops++
			}
		}
	}
	return list.Iterate()
}

func (m *EdgeWeightedGraph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d %d\n", m.v, m.e))
	for v := 0; v < m.v; v++ {
		s.WriteString(fmt.Sprintf("%d : ", v))
		vAdj := m.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			s.WriteString(e.(*Edge).String() + " ")
		}
		s.WriteString("\n")
	}
	return s.String()
}
