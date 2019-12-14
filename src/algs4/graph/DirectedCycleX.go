package graph

import (
	"algs4/queue"
	"algs4/stack"
	"util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */
type DirectedCycleX struct {
	cycle *stack.Stack
}

func NewDirectedCycleX(g *digraph) *DirectedCycle {
	c := &DirectedCycle{}

	// indegrees of remaining vertices
	indegree := make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		indegree[v] = g.Indegree(v)
	}

	// initialize queue to contain all vertices with indegree = 0
	q := queue.NewQueue()
	for v, in := range indegree {
		if in == 0 {
			q.Enqueue(v)
		}
	}
	for !q.IsEmpty() {
		v := q.Dequeue().(int)
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			indegree[w.(int)]--
			if indegree[w.(int)] == 0 {
				q.Enqueue(w.(int))
			}
		}
	}

	// indegree>0 :cycle
	root := -1
	edgeTo := make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if indegree[v] == 0 {
			continue
		}
		root = v
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			edgeTo[w.(int)] = v
		}
	}

	// find any vertex on cycle
	if root != -1 {
		c.cycle = stack.NewStack()
		v := edgeTo[root]
		for x := v; x != root; x = edgeTo[x] {
			c.cycle.Push(x)
		}
		c.cycle.Push(root)
		c.cycle.Push(v)
	}
	return c
}

func (c *DirectedCycleX) HasCycle() bool { return c.cycle != nil }

func (c *DirectedCycleX) Cycle() util.Iterators {
	if !c.HasCycle() {
		return nil
	}
	return c.cycle.Iterate()
}

// todo: check
