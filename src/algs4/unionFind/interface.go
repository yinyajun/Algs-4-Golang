package unionFind

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
	union(p int, q int)
}
