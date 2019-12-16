package graph

import (
	"fmt"
	"util"
)

/**
*
* Immutable weighted Edge
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Edge struct {
	v      int
	w      int
	weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	e := &Edge{}
	if v < 0 {
		panic("NewEdge:invalid vertex")
	}
	if w < 0 {
		panic("NewEdge:invalid vertex")
	}
	e.v = v
	e.w = w
	e.weight = weight
	return e
}

func (e *Edge) Weight() float64 { return e.weight }

func (e *Edge) Either() int { return e.v }

func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	} else if vertex == e.w {
		return e.v
	} else {
		panic("Other: invalid vertex")
	}
}

func (e *Edge) CompareTo(that *Edge) bool {
	return util.Less(e.weight, that.weight)
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}

type EdgeComparator struct{}

func (c EdgeComparator) Compare(i, j interface{}) int {
	if i.(*Edge).Weight() > j.(*Edge).Weight() {
		return 1
	} else if i.(*Edge).Weight() == j.(*Edge).Weight() {
		return 0
	}
	return -1
}
