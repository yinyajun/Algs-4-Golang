package graph

import (
	"strings"
	"fmt"

	. "algs4/bag"
	. "util"
)

/**
* undirected graph
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type graph struct {
	v   int    // # vertex
	e   int    // # edge
	adj []*Bag // Adjacency list
}

func NewGraph(V int) *graph {
	if V < 0 {
		panic("newGraph: invalid v")
	}
	g := &graph{}
	g.v = V
	g.adj = make([]*Bag, V)
	// initialize pointer
	for idx := range g.adj {
		g.adj[idx] = &Bag{}
	}
	return g
}

func NewGraphWithIn(in *In) *graph {
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

func (g *graph) V() int { return g.v }

func (g *graph) E() int { return g.e }

func (g *graph) validateVertex(v int) {
	if v < 0 || v >= g.v {
		panic("validateVertex: invalid vertex")
	}
}

// Adds the undirected edge v-w to this graph.
func (g *graph) AddEdge(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

// Returns the vertices adjacent to vertex v
func (g *graph) Adj(v int) *Bag {
	g.validateVertex(v)
	return g.adj[v]
}

// Returns the degree of vertex v
func (g *graph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

//func (g *graph) String() string {
//	s := strings.Builder{}
//	s.WriteString(fmt.Sprintf("%d vertices, %d edges \n", g.v, g.e))
//	for v := 0; v < g.v; v++ {
//		s.WriteString(fmt.Sprintf("%d:", v))
//		for _, w := range g.adj[v].Iterator() {
//			s.WriteString(fmt.Sprintf("%d ", w.(int)))
//		}
//		s.WriteString("\n")
//	}
//	return s.String()
//}

func (g *graph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d vertices, %d edges \n", g.v, g.e))
	for v := 0; v < g.v; v++ {
		s.WriteString(fmt.Sprintf("%d:", v))
		generator := g.adj[v].Yield()
		for hasNext, w := generator(); hasNext; hasNext, w = generator() {
			s.WriteString(fmt.Sprintf("%d ", w.(int)))
		}
		s.WriteString("\n")
	}
	return s.String()
}
