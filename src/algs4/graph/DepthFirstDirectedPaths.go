package graph

import (
	"util"
	"algs4/stack"
)

/**
* DepthFirstDirectedPaths: finding directed paths from
* a source vertex s to every other vertex in the digraph.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DepthFirstDirectedPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDepthFirstDirectedPaths(g *digraph, s int) *DepthFirstDirectedPaths {
	m := &DepthFirstDirectedPaths{}
	m.marked = make([]bool, g.V())
	m.edgeTo = make([]int, g.V())
	m.validateVertex(s)
	m.s = s
	m.dfs(g, s)
	return m
}

func (m *DepthFirstDirectedPaths) validateVertex(v int) {
	V := len(m.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *DepthFirstDirectedPaths) dfs(g *digraph, v int) {
	m.marked[v] = true

	vAdj := g.Adj(v)
	for hasNext, w := vAdj(); hasNext; hasNext, w = vAdj() {
		if !m.marked[w.(int)] {
			m.edgeTo[w.(int)] = v
			m.dfs(g, w.(int))
		}
	}
}

func (m *DepthFirstDirectedPaths) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.marked[v]
}

func (m *DepthFirstDirectedPaths) PathTo(v int) util.Generator {
	m.validateVertex(v)
	path := stack.NewStack()

	for x := v; x != m.s; x = m.edgeTo[x] {
		path.Push(x)
	}
	path.Push(m.s)
	return path.Yield()
}
