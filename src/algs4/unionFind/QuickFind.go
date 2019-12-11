package unionFind

/**
* QuickFind
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

func (m *QuickFindUF) Count() int {
	return m.cnt
}

func (m *QuickFindUF) Connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *QuickFindUF) find(p int) int {
	return m.id[p]
}

func (m *QuickFindUF) Union(p int, q int) {
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
