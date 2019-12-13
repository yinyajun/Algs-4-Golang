package graph

import "algs4/queue"

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
	return m
}

// todo: edge weight digraph

func (m *DepthFirstOrder) dfs(g *digraph, v int) {
	m.marked[v]=true
	m.pre[v]=m.preCounter
	m.preCounter++
	m.preorder.Enqueue(v)
	vAdj := g.Adj(v)
	for ok, w:= vAdj(); ok; ok,w=vAdj(){
		if !m.marked[w.(int)]{
			m.dfs(g, w.(int))
		}
	}
	m.post[v] = m.postCounter
	m.postCounter++
	m.postorder.Enqueue(v)
}

