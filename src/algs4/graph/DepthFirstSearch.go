package graph

type DepthFirstSearch struct {
	marked []bool // marked[v] = is there an s-v path?
	count  int    // number of vertices connected to s
}

func NewDFS(g *graph, s int) *DepthFirstSearch {
	m := &DepthFirstSearch{make([]bool, g.V()), 0}
	m.validateVertex(s)
	m.dfs(g, s)
	return m
}

func (m *DepthFirstSearch) validateVertex(v int) {
	V := len(m.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *DepthFirstSearch) dfs(g *graph, v int) {
	m.count++
	m.marked[v] = true
	for _, w := range g.Adj(v).Iterator() {
		if !m.marked[w.(int)] {
			m.dfs(g, w.(int))
		}
	}
}

func (m *DepthFirstSearch) Marked(v int) bool {
	m.validateVertex(v)
	return m.marked[v]
}

func (m *DepthFirstSearch) Count() int {
	return m.count
}

func (m *DepthFirstSearch) Search(g *graph, s int) {
	m.dfs(g, s)
}
