package graph

import (
	"util"
	"fmt"
)

/**
*
* Immutable weighted edge
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type edge struct {
	v      int
	w      int
	weight float64
}

func NewEdge(v, w int, weight float64) *edge {
	e := &edge{}
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

func (e *edge) Weight() float64 { return e.weight }

func (e *edge) Either() int { return e.v }

func (e *edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	} else if vertex == e.w {
		return e.v
	} else {
		panic("Other: invalid vertex")
	}
}

func (e *edge) CompareTo(that *edge) bool {
	return util.Less(e.weight, that.weight)
}

func (e *edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}
