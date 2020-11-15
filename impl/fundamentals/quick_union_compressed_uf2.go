/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 15:38
 */

package fundamentals

type QuickUnionCompressedUF2 struct {
	*QuickUnionUF
	rank []int
}

func NewQuickUnionCompressed2(n int) *QuickUnionCompressedUF2 {
	base := &QuickUnionUF{
		parent: make([]int, n),
		count:  n,
	}
	uf := &QuickUnionCompressedUF2{}
	uf.QuickUnionUF = base
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (u *QuickUnionCompressedUF2) Find(p int) int {
	u.validate(p)
	return u.find(p)
}

func (u *QuickUnionCompressedUF2) find(p int) int {
	if u.parent[p] != p {
		u.parent[p] = u.find(u.parent[p]) // non-root node link to the root node
	}
	return u.parent[p] //attention
}

func (u *QuickUnionCompressedUF2) Connected(p, q int) bool { return u.Find(p) == u.Find(q) }

func (u *QuickUnionCompressedUF2) Union(p, q int) {
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
