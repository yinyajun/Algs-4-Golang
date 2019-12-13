package graph

import (
	"algs4/queue"
	"algs4/stack"
	"fmt"
	. "util"
)

/**
* breadth first paths
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

const INT_MAX = int(^uint(0) >> 1)

type BreadthFirstPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	distTo []int  // distTo[v] = number of edges shortest s-v path
	s      int    // source vertex
}

// Computes the shortest path between the source vertex s
// and every other vertex in the graph g.
func NewBreadthFirstPaths(g *graph, s int) *BreadthFirstPaths {
	m := &BreadthFirstPaths{}
	m.s = s
	m.edgeTo = make([]int, g.V())
	m.distTo = make([]int, g.V())
	m.marked = make([]bool, g.V())
	m.validateVertex(s)
	m.bfs(g, s)

	if !m.check(g, s) {
		panic("check: check error")
	}
	return m
}

// Computes the shortest path between any one of the source vertices in sources
// and every other vertex in graph g.
func NewBreadthFirstPathsMultiSources(g *graph, sources []int) *BreadthFirstPaths {
	m := &BreadthFirstPaths{}
	m.edgeTo = make([]int, g.V())
	m.distTo = make([]int, g.V())
	m.marked = make([]bool, g.V())

	for v := 0; v < g.V(); v++ {
		m.distTo[v] = INT_MAX
	}
	m.validateVertices(sources)
	m.bfsMultiSources(g, sources)
	return m
}

func (m *BreadthFirstPaths) validateVertex(v int) {
	V := len(m.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *BreadthFirstPaths) validateVertices(vertices []int) {
	if len(vertices) == 0 {
		panic("validateVertices: empty vertices")
	}
	V := len(m.marked)
	for _, v := range vertices {
		if v < 0 || v >= V {
			panic("validateVertex: invalid vertex")
		}
	}
}

// breadth-first search from a single source
func (m *BreadthFirstPaths) bfs(g *graph, s int) {
	q := queue.NewQueue()
	for idx := range m.distTo {
		m.distTo[idx] = INT_MAX
	}
	m.distTo[s] = 0
	m.marked[s] = true
	q.Enqueue(s)

	for !q.IsEmpty() {
		v := q.Dequeue()
		gen := g.Adj(v.(int))
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			if !m.marked[w.(int)] {
				m.edgeTo[w.(int)] = v.(int)
				m.distTo[w.(int)] = m.distTo[v.(int)] + 1
				m.marked[w.(int)] = true
				q.Enqueue(w)
			}
		}
	}
}

// breadth-first search from multiple sources
func (m *BreadthFirstPaths) bfsMultiSources(g *graph, sources []int) {
	q := queue.NewQueue()
	for _, s := range sources {
		m.marked[s] = true
		m.distTo[s] = 0
		q.Enqueue(s)
	}
	for !q.IsEmpty() {
		v := q.Dequeue().(int)
		vAdj := g.Adj(v)
		for hasNext, w := vAdj(); hasNext; hasNext, w = vAdj() {
			if !m.marked[w.(int)] {
				m.edgeTo[w.(int)] = v
				m.marked[w.(int)] = true
				m.distTo[w.(int)] = m.distTo[v] + 1
				q.Enqueue(w)
			}
		}
	}
}

func (m *BreadthFirstPaths) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.marked[v]
}

func (m *BreadthFirstPaths) PathTo(v int) Generator {
	m.validateVertex(v)
	if !m.marked[v] {
		return nil
	}
	path := stack.NewStack()
	for x := v; x != m.s; x = m.edgeTo[x] {
		path.Push(x)
	}
	path.Push(m.s)
	return path.Yield()
}

func (m *BreadthFirstPaths) DistTo(v int) int {
	m.validateVertex(v)
	return m.distTo[v]
}

// check optimality conditions for single source
func (m *BreadthFirstPaths) check(g *graph, s int) bool {
	// check that the distance of s = 0
	if m.distTo[s] != 0 {
		fmt.Println("distance of source", s, "to itself=", m.distTo[s])
		return false
	}

	// check that for each edge v-w dist[w] <= dist[v] + 1
	// provided v is reachable from s
	for v := 0; v < g.V(); v++ {
		gen := g.Adj(v)
		for hasNext, w := gen(); hasNext; hasNext, w = gen() {
			if m.HasPathTo(v) != m.HasPathTo(w.(int)) {
				fmt.Println("edge", v, "-", w.(int))
				fmt.Println("hasPathTo(", v, ") =", m.HasPathTo(v))
				fmt.Println("hasPathTo(", w.(int), ") =", m.HasPathTo(w.(int)))
				return false
			}
			if m.HasPathTo(v) && m.distTo[w.(int)] > m.distTo[v]+1 {
				fmt.Println("edge", v, "-", w.(int))
				fmt.Println("distTo[", v, "] =", m.DistTo(v))
				fmt.Println("distTo[", w.(int), "] =", m.DistTo(w.(int)))
				return false
			}
		}
	}

	// check that v = edgeTo[w] satisfies distTo[w] = distTo[v] + 1
	// provided v is reachable from s
	for w := 0; w < g.V(); w++ {
		if !m.HasPathTo(w) || w == m.s {
			continue
		}
		v := m.edgeTo[w]
		if m.distTo[w] != m.distTo[v]+1 {
			fmt.Println("shortest path edge", v, "-", w)
			fmt.Println("distTo[", v, "] =", m.DistTo(v))
			fmt.Println("distTo[", w, "] =", m.DistTo(w))
			return false
		}
	}
	return true
}
