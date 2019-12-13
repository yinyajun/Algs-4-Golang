package graph

import (
	. "algs4/stack"
	"util"
)

/**
* check cycle in a undirected graph
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Cycle struct {
	marked   []bool
	edgeTo   []int
	hasCycle bool
	cycle    *Stack
}

/**
 * Determines whether the undirected graph {@code G} has a cycle and,
 * if so, finds such a cycle.
 */
func NewCycle(g *graph) *Cycle {
	c := &Cycle{}
	if c.hasSelfLoop(g) {
		return c
	}
	if c.hasParallelEdges(g) {
		return c
	}
	c.marked = make([]bool, g.V())
	c.edgeTo = make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if !c.marked[v] {
			c.dfs(g, v, v)
		}
	}
	return c
}

// does this graph have a self loop?
// side effect: initialize cycle to be self loop
func (c *Cycle) hasSelfLoop(g *graph) bool {
	for v := 0; v < g.V(); v++ {
		gen := g.Adj(v)
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			if v == w.(int) {
				c.cycle = NewStack()
				c.cycle.Push(v)
				c.cycle.Push(w.(int))
				return true
			}
		}
	}
	return false
}

// does this graph have two parallel edges?
// side effect: initialize cycle to be two parallel edges
func (c *Cycle) hasParallelEdges(g *graph) bool {
	for v := 0; v < g.V(); v++ {
		marked := make([]bool, g.V())
		gen := g.Adj(v)
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			if marked[w.(int)] {
				c.cycle = NewStack()
				c.cycle.Push(v)
				c.cycle.Push(w)
				c.cycle.Push(v)
				return true
			}
			marked[w.(int)] = true
		}
		// reset so marked[v] = false for all v
		gen = g.Adj(v)
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			marked[w.(int)] = false
		}
	}
	return false
}

func (c *Cycle) dfs(g *graph, v int, u int) {
	c.marked[v] = true

	gen := g.Adj(v)
	for hasNext, w := gen(); hasNext; hasNext, w = gen() {
		// short circuit if cycle already found
		if c.HasCycle() {
			return
		}
		if !c.marked[w.(int)] {
			c.edgeTo[w.(int)] = v
			c.dfs(g, w.(int), v)
		} else if w.(int) != u {
			// check for cycle (but disregard reverse of edge leading to v)
			c.cycle = NewStack()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}
}

func (c *Cycle) HasCycle() bool { return c.cycle != nil }

func (c *Cycle) Cycle() util.Generator {
	if !c.HasCycle() {
		return nil
	}
	return c.cycle.Yield()
}
