package graph

import (
	"algs4/stack"
	"fmt"
	"util"
)
/**
* directed cycle
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */


type DirectedCycle struct {
	marked  []bool       // marked[v] = has vertex v been marked?
	edgeTo  []int        // edgeTo[v] = previous vertex on path to v
	onStack []bool       // onStack[v] = is vertex on the stack? 递归调用期间栈上的所有顶点
	cycle   *stack.Stack // directed cycle (or null if no such cycle)
}

// Determines whether the digraph g has a directed cycle and, if so,
// finds such a cycle.
func NewDirectedCycle(g *digraph) *DirectedCycle {
	m := &DirectedCycle{}
	m.marked = make([]bool, g.V())
	m.onStack = make([]bool, g.V())
	m.edgeTo = make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if !m.marked[v] {
			m.dfs(g, v)
		}
	}
	return m
}

// a little different from find cycle in undirected graph,
func (m *DirectedCycle) dfs(g *digraph, v int) {
	m.marked[v] = true
	m.onStack[v] = true

	vAdj := g.Adj(v)
	for ok, w := vAdj(); ok; ok, w = vAdj() {
		// short circuit if directed cycle found
		if m.HasCycle() {
			return
		}
		// found new vertex, so recur
		if !m.marked[w.(int)] {
			m.edgeTo[w.(int)] = v
			m.dfs(g, w.(int))
		} else if m.onStack[w.(int)] {
			m.cycle = stack.NewStack()
			for x := v; x != w; x = m.edgeTo[x] {
				m.cycle.Push(x)
			}
			m.cycle.Push(w)
			m.cycle.Push(v)
			if !m.check() {
				panic("dfs: check failed")
			}
		}
	}
	m.onStack[v] = false // backtracking
}

func (m *DirectedCycle) HasCycle() bool { return m.cycle != nil }

func (m *DirectedCycle) Cycle() util.Generator {
	if !m.HasCycle() {
		return nil
	}
	return m.cycle.Yield()
}

// certify that digraph has a directed cycle if it reports one
func (m *DirectedCycle) check() bool {
	if m.HasCycle() {
		first, last := -1, -1
		gen := m.Cycle()
		for ok, v := gen(); ok; ok, v = gen() {
			if first == -1 {
				first = v.(int)
			}
			last = v.(int)
		}

		if first != last {
			fmt.Printf("cycle begins with %d and ends with %d\n", first, last)
			return false
		}
	}
	return true
}