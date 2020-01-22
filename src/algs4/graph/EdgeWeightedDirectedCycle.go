package graph

import (
	"algs4/stack"
	"util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type EdgeWeightedDirectedCycle struct {
	marked  []bool          // marked[v] = has vertex v been marked?
	edgeTo  []*DirectedEdge // edgeTo[v] = previous edge on path to v
	onStack []bool          // onStack[v] = is vertex on the stack?
	cycle   *stack.Stack    // directed cycle (or null if no such cycle)
}

func NewEdgeWeightedDirectedCycle(g *EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	m := &EdgeWeightedDirectedCycle{}
	m.marked = make([]bool, g.V())
	m.edgeTo = make([]*DirectedEdge, g.V())
	m.onStack = make([]bool, g.V())
	for v := 0; v < g.V(); v++ {
		if !m.marked[v] {
			m.dfs(g, v)
		}
	}

	// check that digraph has a cycle
	if !m.check() {
		panic("check failed")
	}
	return m
}

// check that algorithm computes either the topological order or finds a directed cycle
func (m *EdgeWeightedDirectedCycle) dfs(g *EdgeWeightedDigraph, v int) {
	m.marked[v] = true
	m.onStack[v] = true
	vAdjEdges := g.Adj(v)
	for e := vAdjEdges.Next(); e != nil; e = vAdjEdges.Next() {
		w := e.(*DirectedEdge).To()
		// short circuit if directed cycle found
		if m.cycle != nil {
			return
		}
		// found new vertex, so recur
		if !m.marked[w] {
			m.edgeTo[w] = e.(*DirectedEdge)
			m.dfs(g, w)
		} else if m.onStack[w] { // trace back directed cycle
			m.cycle = stack.NewStack()
			f := e.(*DirectedEdge)
			for f.From() != w {
				m.cycle.Push(f)
				f = m.edgeTo[f.From()]
			}
			m.cycle.Push(f)
			return
		}
	}
	m.onStack[v] = false
}

func (m *EdgeWeightedDirectedCycle) HasCycle() bool { return m.cycle != nil }

func (m *EdgeWeightedDirectedCycle) Cycle() util.Iterator {
	if !m.HasCycle() {
		return nil
	}
	return m.cycle.Iterate()
}

func (m *EdgeWeightedDirectedCycle) check() bool {

	return true
}
