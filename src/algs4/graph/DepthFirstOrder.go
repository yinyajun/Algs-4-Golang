package graph

import (
	"algs4/queue"
	"algs4/stack"
	"fmt"
	"util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    *queue.Queue
	postorder   *queue.Queue
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	m := &DepthFirstOrder{}
	m.marked = make([]bool, g.V())
	m.pre = make([]int, g.V())
	m.post = make([]int, g.V())
	m.preorder = queue.NewQueue()
	m.postorder = queue.NewQueue()
	for v := 0; v < g.V(); v++ {
		if !m.marked[v] {
			m.dfs(g, v)
		}
	}
	if !m.check() {
		panic("check failed")
	}
	return m
}

func NewDepthFirstOrderEWD(g *EdgeWeightedDigraph) *DepthFirstOrder {
	m := &DepthFirstOrder{}
	m.marked = make([]bool, g.V())
	m.pre = make([]int, g.V())
	m.post = make([]int, g.V())
	m.preorder = queue.NewQueue()
	m.postorder = queue.NewQueue()
	for v := 0; v < g.V(); v++ {
		if !m.marked[v] {
			m.dfsEWD(g, v)
		}
	}
	return m
}

func (m *DepthFirstOrder) dfs(g *Digraph, v int) {
	m.marked[v] = true
	m.pre[v] = m.preCounter
	m.preCounter++
	m.preorder.Enqueue(v)
	vAdj := g.Adj(v)
	for w := vAdj.Next(); w != nil; w = vAdj.Next() {
		if !m.marked[w.(int)] {
			m.dfs(g, w.(int))
		}
	}
	m.post[v] = m.postCounter
	m.postCounter++
	m.postorder.Enqueue(v)
}

func (m *DepthFirstOrder) dfsEWD(g *EdgeWeightedDigraph, v int) {
	m.marked[v] = true
	m.pre[v] = m.preCounter
	m.preCounter++
	m.preorder.Enqueue(v)
	vAdjEdges := g.Adj(v)
	for e := vAdjEdges.Next(); e != nil; e = vAdjEdges.Next() {
		w := e.(*DirectedEdge).To()
		if !m.marked[w] {
			m.dfsEWD(g, w)
		}
	}
	m.post[v] = m.postCounter
	m.postCounter++
	m.postorder.Enqueue(v)
}

func (m *DepthFirstOrder) Pre(v int) int {
	m.validateVertex(v)
	return m.pre[v]
}
func (m *DepthFirstOrder) Post(v int) int {
	m.validateVertex(v)
	return m.post[v]
}

func (m *DepthFirstOrder) PostOrder() util.Iterator { return m.postorder.Iterate() }

func (m *DepthFirstOrder) PreOrder() util.Iterator { return m.preorder.Iterate() }

func (m *DepthFirstOrder) ReversePostOrder() util.Iterator {
	reverse := stack.NewStack()
	gen := m.PostOrder()
	for v := gen.Next(); v != nil; v = gen.Next() {
		reverse.Push(v)
	}
	return reverse.Iterate()
}

func (m *DepthFirstOrder) validateVertex(v int) {
	V := len(m.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *DepthFirstOrder) check() bool {
	r := 0
	order := m.PreOrder()
	for v := order.Next(); v != nil; v = order.Next() {
		if m.Pre(v.(int)) != r {
			fmt.Println("pre not consistent")
			return false
		}
		r++
	}

	r = 0
	order = m.PostOrder()
	for v := order.Next(); v != nil; v = order.Next() {
		if m.Post(v.(int)) != r {
			fmt.Println("post not consistent")
			return false
		}
		r++
	}
	return true
}
