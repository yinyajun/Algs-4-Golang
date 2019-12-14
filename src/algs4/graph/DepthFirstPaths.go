package graph

import (
	"algs4/stack"
	. "util"
)

/**
* depth first paths
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DepthFirstPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	s      int    // source vertex
}

func NewDepthFirstPaths(g *graph, s int) *DepthFirstPaths {
	m := &DepthFirstPaths{}
	m.s = s
	m.edgeTo = make([]int, g.V())
	m.marked = make([]bool, g.V())
	m.validateVertex(s)
	m.dfs(g, s)
	return m
}

func (m *DepthFirstPaths) validateVertex(v int) {
	V := len(m.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

// depth first search from v
func (m *DepthFirstPaths) dfs(g *graph, v int) {
	m.marked[v] = true
	vAdj := g.Adj(v)
	for w := vAdj.Next(); w != nil; w = vAdj.Next() {
		if !m.marked[w.(int)] {
			m.edgeTo[w.(int)] = v
			m.dfs(g, w.(int))
		}
	}
}

// Is there a path between the source vertex s and vertex v?
func (m *DepthFirstPaths) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.marked[v]
}

func (m *DepthFirstPaths) PathTo(v int) Iterator {
	m.validateVertex(v)
	if !m.HasPathTo(v) {
		return nil
	}
	path := stack.NewStack()
	for x := v; x != m.s; x = m.edgeTo[x] {
		path.Push(x)
	}
	path.Push(m.s)
	return path.Iterate()
}
