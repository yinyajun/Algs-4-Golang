package graph

import (
	"fmt"
	"strings"

	. "algs4/bag"
	. "util"
)

/**
* undirected Graph
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Graph struct {
	v   int    // # vertex
	e   int    // # Edge
	adj []*Bag // Adjacency list
}

func NewGraph(V int) *Graph {
	if V < 0 {
		panic("newGraph: invalid v")
	}
	g := &Graph{}
	g.v = V
	g.adj = make([]*Bag, V)
	// initialize pointer
	for idx := range g.adj {
		g.adj[idx] = &Bag{}
	}
	return g
}

func NewGraphWithIn(in *In) *Graph {
	v := in.ReadInt()
	g := NewGraph(v)
	e := in.ReadInt()
	if e < 0 {
		panic("NewGraphWithIn: invalid e")
	}
	for i := 0; i < e; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		g.AddEdge(v, w)
	}
	return g
}

func (g *Graph) V() int { return g.v }

func (g *Graph) E() int { return g.e }

func (g *Graph) validateVertex(v int) {
	if v < 0 || v >= g.v {
		panic("validateVertex: invalid vertex")
	}
}

// Adds the undirected Edge v-w to this Graph.
func (g *Graph) AddEdge(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

// Returns the vertices adjacent to vertex v
func (g *Graph) Adj(v int) Iterator {
	g.validateVertex(v)
	return g.adj[v].Iterate()
}

// Returns the degree of vertex v
func (g *Graph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

func (g *Graph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d vertices, %d edges \n", g.v, g.e))
	for v := 0; v < g.v; v++ {
		s.WriteString(fmt.Sprintf("%d:", v))
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			s.WriteString(fmt.Sprintf("%d ", w.(int)))
		}
		s.WriteString("\n")
	}
	return s.String()
}
