package graph

import . "algs4/priorityQueue"

/**
* a data type for solving the single-source shortest paths problem in edge-weighted digraphs
* where the edge weights are nonnegative.
*
* This implementation  uses Dijkstra's algorithm with a binary heap.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DijkstraSP struct {
	distTo []float64       // distTo[v] = distance  of shortest s->v path
	edgeTo []*DirectedEdge // edgeTo[v] = last edge on shortest s->v path
	pq     *IndexMinPQ     // priority queue of vertices
}

func NewDijkstraSP(g *EdgeWeightedDigraph, s int) *DijkstraSP {
	m := &DijkstraSP{}
	edges := g.Edges()
	for e := edges.Next(); e != nil; e = edges.Next() {
		if e.(*DirectedEdge).Weight() < 0 {
			panic("NewDijkstraSP: edge has negative weight")
		}
	}
	m.validateVertex(s)

	m.distTo = make([]float64, g.V())
	m.edgeTo = make([]*DirectedEdge, g.V())
	for v := 0; v < g.V(); v++ {
		m.distTo[v] = POSTIVE_INFINITY
	}
	m.pq = NewIndexMinPQ(g.V())

	m.distTo[s] = 0.0
	// relax vertices in order of distance from s
	m.pq.Insert(s, m.distTo[s])
	for !m.pq.IsEmpty() {
		v := m.pq.DelMin()
		vAdj := g.Adj(v)
		for e := vAdj.Next(); e != nil; e = vAdj.Next() {
			m.relax(e.(*DirectedEdge))
		}
	}

	// check optimality conditions
	if !m.check(g, s) {
		panic("check failed")
	}
	return m
}

func (m *DijkstraSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if m.distTo[v]+e.Weight() < m.distTo[w] {
		m.distTo[w] = m.distTo[v] + e.Weight()
		m.edgeTo[w] = e
		if m.pq.Contains(w) {
			m.pq.DecreaseKey(w, m.distTo[w])
		} else {
			m.pq.Insert(w, m.distTo[w])
		}
	}
}


func (m *DijkstraSP) DistTo(v int) float64{
	m.validateVertex(v)
	return m.distTo[v]
}


func (m*DijkstraSP) HasPathTo(v int) bool{
	m.validateVertex(v)
	return m.distTo[v] < POSTIVE_INFINITY
}



