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

type QuickFindUF struct {
	id  []int
	cnt int
}

func NewQuickFindUF(N int) *QuickFindUF {
	id := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	return &QuickFindUF{
		id:  id,
		cnt: N,
	}
}

func (m *QuickFindUF) count() int {
	return m.cnt
}

func (m *QuickFindUF) connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *QuickFindUF) find(p int) int {
	return m.id[p]
}

func (m *QuickFindUF) union(p int, q int) {
	pID := m.find(p)
	qID := m.find(q)
	if pID == qID {
		return
	}
	for idx, v := range m.id {
		if v == pID {
			m.id[idx] = qID
		}
	}
	m.cnt--
}

func EgQuickFindUF() {
	in := io.NewIn(io.SplitFunc("words"))
	N := in.ReadInt()
	uf := NewQuickFindUF(N)
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
