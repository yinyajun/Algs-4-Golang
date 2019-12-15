package graph

import (
	. "algs4/bag"
	"algs4/stack"
	"fmt"
	"strings"
	. "util"
)

/**
*
* Digraph
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Digraph struct {
	v        int    // number of vertices in this Digraph
	e        int    // number of edges in this Digraph
	adj      []*Bag // adj[v] = adjacency list for vertex v
	indegree []int  // indegree[v] = indegree of vertex v
}

// Initializes an empty Digraph with V vertices.
func NewDigraph(V int) *Digraph {
	if V < 0 {
		panic("NewDigraph: invalid V")
	}
	dg := &Digraph{}
	dg.v = V
	dg.adj = make([]*Bag, V)
	for idx := range dg.adj {
		dg.adj[idx] = &Bag{}
	}
	dg.indegree = make([]int, V)
	return dg
}

// Initializes a Digraph from the specified input stream.
func NewDigraphWithIn(in *In) *Digraph {
	v := in.ReadInt()
	dg := NewDigraph(v)
	e := in.ReadInt()
	if e < 0 {
		panic("NewDigraphWithIn: invalid e")
	}
	for i := 0; i < e; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		dg.AddEdge(v, w)
	}
	return dg
}

// Initializes a new Digraph that is a deep copy of the specified Digraph.
func NewDigraphWithGraph(g *Digraph) *Digraph {
	dg := NewDigraph(g.V())
	dg.e = g.E()
	for v := 0; v < g.V(); v++ {
		dg.indegree[v] = g.Indegree(v)
	}
	for v := 0; v < g.V(); v++ {
		// reverse so that adjacency list is in same order as original
		reverse := stack.NewStack()
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			reverse.Push(w)
		}
		it := reverse.Iterate()
		for w := it.Next(); w != nil; w = it.Next() {
			dg.adj[v].Add(w)
		}
	}
	return dg
}

func (dg *Digraph) V() int { return dg.v }

func (dg *Digraph) E() int { return dg.e }

func (dg *Digraph) validateVertex(v int) {
	if v < 0 || v >= dg.v {
		panic("validateVertex: invalid vertex")
	}
}

// Adds the directed Edge v→w to this Digraph.
func (dg *Digraph) AddEdge(v, w int) {
	dg.validateVertex(v)
	dg.validateVertex(w)
	dg.adj[v].Add(w)
	dg.indegree[w]++
	dg.e++
}

func (dg *Digraph) Adj(v int) Iterator {
	dg.validateVertex(v)
	return dg.adj[v].Iterate()
}

func (dg *Digraph) Reverse() *Digraph {
	reverse := NewDigraph(dg.v)
	for v := 0; v < dg.v; v++ {
		vAdj := dg.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			reverse.adj[w.(int)].Add(v)
		}
	}
	return reverse
}

func (dg *Digraph) Indegree(v int) int {
	dg.validateVertex(v)
	return dg.indegree[v]
}

func (dg *Digraph) Outdegree(v int) int {
	dg.validateVertex(v)
	return dg.adj[v].Size()
}

func (dg *Digraph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d vertices, %d edges \n", dg.v, dg.e))
	for v := 0; v < dg.V(); v++ {
		s.WriteString(fmt.Sprintf("%d: ", v))
		vAdj := dg.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			s.WriteString(fmt.Sprintf("%d ", w))
		}
		s.WriteString("\n")
	}
	return s.String()
}
