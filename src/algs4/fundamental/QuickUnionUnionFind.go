package fundamental

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type QuickFindUF struct {
	UFBase
}

func NewQuickUnionUF(N int) *QuickFindUF {
	id := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	return &QuickFindUF{
		UFBase: UFBase{
			id:  id,
			cnt: N,
		},
	}
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
