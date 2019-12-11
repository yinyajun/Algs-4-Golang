package unionFind

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type UnionFind interface {
	Count() int
	Connected(int, int) bool
	find(int) int
	Union(p int, q int)
}
