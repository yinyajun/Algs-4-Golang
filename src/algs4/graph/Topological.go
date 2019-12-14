package graph

import "util"

/**
* determining a topological order of a directed acyclic graph (DAG)
* Recall, a digraph has a topological order if and only if it is a DAG.
* This implementation uses depth-first search.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type topological struct {
	order util.Iterator
	rank  []int
}

func NewTopological(g *digraph) *topological {
	t := &topological{}
	t.rank = make([]int, g.V())
	finder := NewDirectedCycle(g)
	if finder.HasCycle() {
		return t
	}
	dfs := NewDepthFirstOrder(g)
	t.order = dfs.ReversePostOrder()
	i := 0
	for v := t.order.Next(); v != nil; v = t.order.Next() {
		t.rank[v.(int)] = i
		i++
	}
	return t
}

// todo: topological edge weight digraph

func (t *topological) Order() util.Iterator {
	t.order.Reset()
	return t.order
}

func (t *topological) HasOrder() bool { return t.order != nil }

func (t *topological) IsDAG() bool { return t.HasOrder() }

func (t *topological) Rank(v int) int {
	t.validateVertex(v)
	if t.HasOrder() {
		return t.rank[v]
	}
	return -1
}

func (t *topological) validateVertex(v int) {
	V := len(t.rank)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
