package graph

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type transitiveClosure struct {
	tc []*DirectedDFS
}

func NewTransitiveClosure(g *digraph) *transitiveClosure {
	m := &transitiveClosure{}
	m.tc = make([]*DirectedDFS, g.V())
	for v := 0; v < g.V(); v++ {
		m.tc[v] = NewDirectedDFS(g, v)
	}
	return m
}

func (m *transitiveClosure) Reachable(v, w int) bool {
	m.validateVertex(v)
	m.validateVertex(w)
	return m.tc[v].Marked(w)
}

func (m *transitiveClosure) validateVertex(v int) {
	V := len(m.tc)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
