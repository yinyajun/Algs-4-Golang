package graph

import (
	"algs4/queue"
	"util"
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

	onQueue []bool        // onQueue[v] = is v currently on the queue?
	queue   *queue.Queue  // queue of vertices to relax
	cost    int           // number of calls to relax()
	cycle   util.Iterator // negative cycle (or nil if no such cycle)
}

func NewBellmanFordSpP(g *EdgeWeightedDigraph, s int) *BellmanFordSP {
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
		m.onQueue[v] = false
		m.relax(g, v)
	}

	if !m.check(g, s) {
		panic("check failed")
	}
	return m
}

// relax vertex v and put other endpoints on queue if changed
func (m *BellmanFordSP) relax(g *EdgeWeightedDigraph, v int) {
	vAdj := g.Adj(v)
	for e := vAdj.Next(); e != nil; e = vAdj.Next() {
		w := e.(*DirectedEdge).To()
		if m.distTo[w] > m.distTo[v]+e.(*DirectedEdge).Weight() {
			m.distTo[w] = m.distTo[w] + e.(*DirectedEdge).Weight()
			m.edgeTo[w] = e.(*DirectedEdge)
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
func (m *BellmanFordSP) findNegativeCycle(){
	v
}