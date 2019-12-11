package unionFind

/**
*	QuickUnion
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

func (m *QuickUnionUF) Count() int {
	return m.cnt
}

func (m *QuickUnionUF) Connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *QuickUnionUF) find(p int) int {
	// 一直找到根节点
	for p != m.id[p] {
		p = m.id[p]
	}
	return p
}

func (m *QuickUnionUF) Union(p int, q int) {
	pRoot := m.find(p)
	qRoot := m.find(q)

	if pRoot == qRoot {
		return
	}
	m.id[pRoot] = qRoot
	m.cnt--
}

