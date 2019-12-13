package graph

import (
	"algs4/queue"
	"algs4/stack"
	"fmt"
	"util"
)

type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    *queue.Queue
	postorder   *queue.Queue
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(g *digraph) *DepthFirstOrder {
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

// todo: edge weight digraph

func (m *DepthFirstOrder) dfs(g *digraph, v int) {
	m.marked[v] = true
	m.pre[v] = m.preCounter
	m.preCounter++
	m.preorder.Enqueue(v)
	vAdj := g.Adj(v)
	for ok, w := vAdj(); ok; ok, w = vAdj() {
		if !m.marked[w.(int)] {
			m.dfs(g, w.(int))
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

func (m *DepthFirstOrder) PostOrder() util.Generator { return m.postorder.Yield() }

func (m *DepthFirstOrder) PreOrder() util.Generator { return m.preorder.Yield() }

func (m *DepthFirstOrder) ReversePostOrder() util.Generator {
	reverse := stack.NewStack()
	gen := m.PostOrder()
	for ok, v := gen(); ok; ok, v = gen() {
		reverse.Push(v)
	}
	return reverse.Yield()
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
	for ok, v := order(); ok; ok, v = order() {
		if m.Pre(v.(int)) != r {
			fmt.Println("pre not consistent")
			return false
		}
		r++
	}

	r = 0
	order = m.PostOrder()
	for ok, v := order(); ok; ok, v = order() {
		if m.Post(v.(int)) != r {
			fmt.Println("post not consistent")
			return false
		}
		r++
	}
	return true
}
