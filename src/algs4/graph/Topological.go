package graph

import "util"

/**
* determining a Topological order of a directed acyclic Graph (DAG)
* Recall, a Digraph has a Topological order if and only if it is a DAG.
* This implementation uses depth-first search.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Topological struct {
	order util.Iterator
	rank  []int
}

func NewTopological(g *Digraph) *Topological {
	t := &Topological{}
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

// todo: Topological Edge weight Digraph

func (t *Topological) Order() util.Iterator {
	t.order.Reset()
	return t.order
}

func (t *Topological) HasOrder() bool { return t.order != nil }

func (t *Topological) IsDAG() bool { return t.HasOrder() }

func (t *Topological) Rank(v int) int {
	t.validateVertex(v)
	if t.HasOrder() {
		return t.rank[v]
	}
	return -1
}

func (t *Topological) validateVertex(v int) {
	V := len(t.rank)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}