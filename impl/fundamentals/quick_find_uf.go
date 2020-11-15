/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 11:43
 */

package fundamentals

import "Algs-4-Golang/utils"

type QuickFindUF struct {
	id    []int
	count int
}

func NewQuickFindUF(n int) *QuickFindUF {
	uf := &QuickFindUF{}
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	uf.id = id
	uf.count = n
	return uf
}

func (u *QuickFindUF) Count() int { return u.count }

func (u *QuickFindUF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *QuickFindUF) Find(p int) int {
	u.validate(p)
	return u.id[p]
}

func (u *QuickFindUF) Union(p, q int) {
	u.validate(p)
	u.validate(q)
	pID := u.id[p]
	qID := u.id[q]

	if pID == qID {
		return
	}

	for i := 0; i < len(u.id); i++ {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}
	u.count--
}

func (u *QuickFindUF) validate(p int) {
	n := len(u.id)
	if p < 0 || p >= n {
		utils.PanicF("index %d is not between 0 and %d", p, n-1)
	}
}
