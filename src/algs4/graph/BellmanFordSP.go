package graph

import (
	"algs4/queue"
	"util"
	"algs4/stack"
	"fmt"
	"math"
)

/**
* a data type for solving the single-source shortest paths problem in edge-weighted digraphs
* with no negative cycles.
*
* This implementation  uses Bellman-Ford-Moore algorithm.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type BellmanFordSP struct {
	distTo []float64       // distTo[v] = distance  of shortest s->v path
	edgeTo []*DirectedEdge // edgeTo[v] = last edge on shortest s->v path

	onQueue []bool       // onQueue[v] = is v currently on the queue?
	queue   *queue.Queue // queue of vertices to relax

	cost  int           // number of calls to relax()
	cycle util.Iterator // negative cycle (or nil if no such cycle)
}

func NewBellmanFordSP(g *EdgeWeightedDigraph, s int) *BellmanFordSP {
	m := &BellmanFordSP{}
	m.distTo = make([]float64, g.V())
	m.edgeTo = make([]*DirectedEdge, g.V())

	m.onQueue = make([]bool, g.V())
	for v := 0; v < g.V(); v++ {
		m.distTo[v] = POSTIVE_INFINITY
	}
	m.distTo[s] = 0.0

	// Bellman-Ford algorithm
	m.queue = queue.NewQueue()
	m.queue.Enqueue(s)
	m.onQueue[s] = true
	for !m.queue.IsEmpty() && !m.HasNegativeCycle() {
		v := m.queue.Dequeue().(int)
		fmt.Println(v)
		m.debug()
		m.onQueue[v] = false
		m.relax(g, v)
	}

	//if !m.check(g, s) {
	//	panic("check failed")
	//}
	return m
}

// relax vertex v and put other endpoints on queue if changed
func (m *BellmanFordSP) relax(g *EdgeWeightedDigraph, v int) {
	vAdj := g.Adj(v)
	for e := vAdj.Next(); e != nil; e = vAdj.Next() {
		w := e.(*DirectedEdge).To()
		fmt.Println(v, w, e.(*DirectedEdge).Weight())
		if m.distTo[w] > m.distTo[v]+e.(*DirectedEdge).Weight() {
			m.distTo[w] = m.distTo[v] + e.(*DirectedEdge).Weight()
			m.edgeTo[w] = e.(*DirectedEdge)
			fmt.Println(v, w, m.distTo[v], m.distTo[w])
			if !m.onQueue[w] {
				m.queue.Enqueue(w)
				m.onQueue[w] = true
			}
		}
		m.cost++
		if m.cost%g.V() == 0 {
			m.findNegativeCycle()
			if m.HasNegativeCycle() {
				return // found a negative cycle
			}
		}
	}
}

func (m *BellmanFordSP) HasNegativeCycle() bool { return m.cycle != nil }

func (m *BellmanFordSP) NegativeCycle() util.Iterator { return m.cycle }

// by finding a cycle in predecessor graph
func (m *BellmanFordSP) findNegativeCycle() {
	V := len(m.edgeTo)
	spt := NewEdgeWeightedDigraph(V)

	for v := 0; v < V; v++ {
		if m.edgeTo[v] != nil {
			spt.AddEdge(m.edgeTo[v])
		}
	}
	finder := NewEdgeWeightedDirectedCycle(spt)
	m.cycle = finder.Cycle()
}

func (m *BellmanFordSP) DistTo(v int) float64 {
	m.validateVertex(v)
	if m.HasNegativeCycle() {
		panic("Negative cost cycle exists.")
	}
	return m.distTo[v]
}

func (m *BellmanFordSP) HasPathTo(v int) bool {
	m.validateVertex(v)
	return m.distTo[v] < POSTIVE_INFINITY
}

func (m *BellmanFordSP) PathTo(v int) util.Iterator {
	m.validateVertex(v)
	if m.HasNegativeCycle() {
		panic("Negative cost cycle exists.")
	}
	if ! m.HasPathTo(v) {
		return nil
	}
	path := stack.NewStack()
	for e := m.edgeTo[v]; e != nil; e = m.edgeTo[e.From()] {
		path.Push(e)
	}
	return path.Iterate()
}

// check optimality conditions: either
// (i) there exists a negative cycle reacheable from s
//     or
// (ii)  for all edges e = v->w:            distTo[w] <= distTo[v] + e.weight()
// (ii') for all edges e = v->w on the SPT: distTo[w] == distTo[v] + e.weight()
func (m *BellmanFordSP) check(g *EdgeWeightedDigraph, s int) bool {

	// has a negative cycle
	if m.HasNegativeCycle() {
		weight := 0.0
		nc := m.NegativeCycle()
		for e := nc.Next(); e != nil; e = nc.Next() {
			weight += e.(*DirectedEdge).Weight()
		}
		if weight >= 0.0 {
			fmt.Println("error: weight of negative cycle = ", weight)
			return false
		}
	} else { // no negative cycle reachable from source
		// check that distTo[v] and edgeTo[v] are consistent
		if m.distTo[s] != 0.0 || m.edgeTo[s] != nil {
			fmt.Println("distTo[s] and edgeTo[s] inconsistent")
			return false
		}
		for v := 0; v < g.V(); v++ {
			if v == s {
				continue
			}
			if m.edgeTo[v] == nil && math.Abs(m.distTo[v]-POSTIVE_INFINITY) > 1E-10 {
				fmt.Println("distTo[] and edgeTo[] inconsistent")
				return false
			}
		}
		// check that all edges e = v->w satisfy distTo[w] <= distTo[v] + e.weight()
		for v := 0; v < g.V(); v++ {
			vAdj := g.Adj(v)
			for e := vAdj.Next(); e != nil; e = vAdj.Next() {
				w := e.(*DirectedEdge).To()
				if m.distTo[v]+e.(*DirectedEdge).Weight() < m.distTo[w] {
					fmt.Printf("edge %s not relaxed", e.(*DirectedEdge))
					return false
				}
			}
		}
		// check that all edges e = v->w on SPT satisfy distTo[w] == distTo[v] + e.weight()
		for w := 0; w < g.V(); w++ {
			if m.edgeTo[w] == nil {
				continue
			}
			e := m.edgeTo[w]
			v := e.From()
			if w != e.To() {
				return false
			}
			if math.Abs(m.distTo[v]+e.Weight()-m.distTo[w]) > 1E-10 {
				fmt.Printf("edge %s on shortest path not tight.", e)
				return false
			}
		}
	}
	fmt.Println("Satisfies optimality conditions")
	fmt.Println()
	return true
}

func (m *BellmanFordSP) validateVertex(v int) {
	V := len(m.distTo)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (m *BellmanFordSP) debug() {
	for i := range m.distTo {
		fmt.Println("\t", i, m.distTo[i])
	}
}
