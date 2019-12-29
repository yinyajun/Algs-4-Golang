package graph

import "fmt"

/**
* A weighted edge in an EdgeWeightedDigraph.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type DirectedEdge struct {
	v      int
	w      int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	m := &DirectedEdge{}
	if v < 0 || w < 0 {
		panic("NewDirectedEdge: Negative Vertex")
	}
	m.v = v
	m.w = w
	m.weight = weight
	return m
}

func (m *DirectedEdge) From() int { return m.v }

func (m *DirectedEdge) To() int { return m.w }

func (m *DirectedEdge) Weight() float64 { return m.weight }

func (m *DirectedEdge) String() string { return fmt.Sprintf("%d -> %d %5.2f", m.v, m.w, m.weight) }
