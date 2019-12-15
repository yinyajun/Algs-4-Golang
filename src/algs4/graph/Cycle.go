package graph

import (
	. "algs4/stack"
	"util"
)

/**
* check cycle in a undirected Graph
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
 * Determines whether the undirected Graph {@code G} has a cycle and,
 * if so, finds such a cycle.
 */
func NewCycle(g *Graph) *Cycle {
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

// does this Graph have a self loop?
// side effect: initialize cycle to be self loop
func (c *Cycle) hasSelfLoop(g *Graph) bool {
	for v := 0; v < g.V(); v++ {
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
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

// does this Graph have two parallel edges?
// side effect: initialize cycle to be two parallel edges
func (c *Cycle) hasParallelEdges(g *Graph) bool {
	for v := 0; v < g.V(); v++ {
		marked := make([]bool, g.V())
		vAdj := g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
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
		vAdj = g.Adj(v)
		for w := vAdj.Next(); w != nil; w = vAdj.Next() {
			marked[w.(int)] = false
		}
	}
	return false
}

func (c *Cycle) dfs(g *Graph, v int, u int) {
	c.marked[v] = true

	vAdj := g.Adj(v)
	for w := vAdj.Next(); w != nil; w = vAdj.Next() {
		// short circuit if cycle already found
		if c.HasCycle() {
			return
		}
		if !c.marked[w.(int)] {
			c.edgeTo[w.(int)] = v
			c.dfs(g, w.(int), v)
		} else if w.(int) != u {
			// check for cycle (but disregard reverse of Edge leading to v)
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

func (c *Cycle) Cycle() util.Iterator {
	if !c.HasCycle() {
		return nil
	}
	return c.cycle.Iterate()
}
