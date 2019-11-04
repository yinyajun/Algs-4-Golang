package fundamental

import (
	"util/io"
	"fmt"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type QuickUnionUF struct {
	id  []int
	cnt int
}

func NewQuickUnionUF(N int) *QuickUnionUF {
	id := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	return &QuickUnionUF{
		id:  id,
		cnt: N,
	}
}

func (m *QuickUnionUF) count() int {
	return m.cnt
}

func (m *QuickUnionUF) connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *QuickUnionUF) find(p int) int {
	for p != m.id[p] {
		p = m.id[p]
	}
	return p
}

func (m *QuickUnionUF) union(p int, q int) {
	pRoot := m.find(p)
	qRoot := m.find(q)

	if pRoot == qRoot {
		return
	}
	m.id[pRoot] = qRoot
	m.cnt--
}

func EgQuickUnionUF() {
	in := io.NewIn(io.SplitFunc("words"))
	N := in.ReadInt()
	uf := NewQuickUnionUF(N)
	for !in.IsEmpty() {
		p := in.ReadInt()
		q := in.ReadInt()
		if uf.connected(p, q) {
			continue
		}
		uf.union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(uf.count(), "components")
}
