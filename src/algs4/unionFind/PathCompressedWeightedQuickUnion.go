package unionFind

/**
*  Path Compressed Weighted Quick Union Union Find
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type PathComWeightedQU struct {
	sz  []int
	id  []int
	cnt int
}

func NewPathComWeightedQU(N int) *PathComWeightedQU {
	id := make([]int, N)
	sz := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	for idx := range sz {
		sz[idx] = 1
	}
	return &PathComWeightedQU{
		id:  id,
		cnt: N,
		sz:  sz,
	}
}

func (m *PathComWeightedQU) Count() int {
	return m.cnt
}

func (m *PathComWeightedQU) Connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *PathComWeightedQU) find(p int) int {
	for p != m.id[p] {
		m.id[p] = m.id[m.id[p]] // 选择父亲的父亲作为自己的新父亲
		p = m.id[p]
	}
	return p
}

func (m *PathComWeightedQU) Union(p int, q int) {
	pRoot := m.find(p)
	qRoot := m.find(q)

	if pRoot == qRoot {
		return
	}
	// 将小树的根节点连接到大树的根节点
	if m.sz[pRoot] < m.sz[qRoot] {
		m.id[pRoot] = qRoot
		m.sz[qRoot] += m.sz[pRoot]
	} else {
		m.id[qRoot] = pRoot
		m.sz[pRoot] += m.sz[qRoot]
	}
	m.cnt--
}
