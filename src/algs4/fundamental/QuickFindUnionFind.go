package fundamental

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type QuickUnionUF struct {
	UFBase
}

func NewQuickFindUF(N int) *QuickUnionUF {
	id := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	return &QuickUnionUF{
		UFBase: UFBase{
			id:  id,
			cnt: N,
		},
	}
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
