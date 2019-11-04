package fundamental

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type UnionFind interface {
	count() int
	connected(int, int) bool
	find(int) int
	union(int int)
}

type UFBase struct {
	id  []int
	cnt int
}

func (m *UFBase) count() int {
	return m.cnt
}

func (m *UFBase) connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *UFBase) find(e int) int {
	return 1
}
