/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 13:23
 */

package fundamentals

import "Algs-4-Golang/utils"

type QuickUnionUF struct {
	parent []int
	count  int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	uf := &QuickUnionUF{}
	uf.count = n
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (u *QuickUnionUF) Find(p int) int {
	u.validate(p)
	for u.parent[p] != p { // p is not root node
		p = u.parent[p]
	}
	// p == u.parent[p], p is root node
	return p
}

func (u *QuickUnionUF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *QuickUnionUF) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	// u.parent[qRoot] = pRoot is also ok
	// 谁都可以当爸爸
	u.parent[pRoot] = qRoot
	u.count--
}

func (u *QuickUnionUF) Count() int { return u.count }

func (u *QuickUnionUF) validate(p int) {
	n := len(u.parent)
	if p < 0 || p >= n {
		utils.PanicF("index %d is not between 0 and %d", p, n-1)
	}
}
