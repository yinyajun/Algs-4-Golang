/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 15:37
 */

package fundamentals

type QuickUnionCompressedUF struct {
	*QuickUnionUF
	rank []int
}

func NewQuickUnionCompressed(n int) *QuickUnionCompressedUF {
	base := &QuickUnionUF{
		parent: make([]int, n),
		count:  n,
	}
	uf := &QuickUnionCompressedUF{}
	uf.QuickUnionUF = base
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

// 查找的过程中，将查找路径压缩
func (u *QuickUnionCompressedUF) Find(p int) int {
	u.validate(p)
	for u.parent[p] != p { // p is not root node
		u.parent[p] = u.parent[u.parent[p]] // path compression by halving
		p = u.parent[p]
	}
	// p == parent[p], p is root node
	return p
}

func (u *QuickUnionCompressedUF) Connected(p, q int) bool { return u.Find(p) == u.Find(q) }

func (u *QuickUnionCompressedUF) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}

	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		//u.rank[pRoot]= u.rank[qRoot]
		u.parent[pRoot] = qRoot
		u.rank[qRoot]++
	}
	u.count--
}
