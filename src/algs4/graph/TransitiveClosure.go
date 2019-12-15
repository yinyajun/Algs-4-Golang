package graph

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type TransitiveClosure struct {
	tc []*DirectedDFS
}

func NewTransitiveClosure(g *Digraph) *TransitiveClosure {
	m := &TransitiveClosure{}
	m.tc = make([]*DirectedDFS, g.V())
	for v := 0; v < g.V(); v++ {
		m.tc[v] = NewDirectedDFS(g, v)
	}
	return m
}

func (m *TransitiveClosure) Reachable(v, w int) bool {
	m.validateVertex(v)
	m.validateVertex(w)
	return m.tc[v].Marked(w)
}

func (m *TransitiveClosure) validateVertex(v int) {
	V := len(m.tc)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
