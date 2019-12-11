package graph

import "algs4/stack"

type DepthFirstPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	s      int    // source vertex
}

func NewDepthFirstPaths(g *graph, s int) {
	m := &DepthFirstPaths{}
	m.s = s
	m.edgeTo = make([]int, g.V())
	m.marked = make([]bool, g.V())
	m.validateVertex(s)
	m.dfs(g, s)
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
	for _, w := range g.Adj(v).Iterator() {
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

func (m *DepthFirstPaths) PathTo(v int) []int {
	ret := []int{}
	m.validateVertex(v)
	if !m.HasPathTo(v) {
		return ret
	}
	path := stack.NewStack()
	for x:= v; x!=m.s; x= m.edgeTo[x]{
		path.Push(x)
	}
	path.Push(m.s)

}
