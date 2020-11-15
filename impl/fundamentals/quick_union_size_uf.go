/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 15:15
 */

package fundamentals

// quick union基于size的优化
// 把size小的根节点链接到size大的根节点，以期望获得更小高度的树
type QuickUnionSizeUF struct {
	*QuickUnionUF
	size []int
}

func NewQuickUnionSizeUF(n int) *QuickUnionSizeUF {
	uf := &QuickUnionSizeUF{
		QuickUnionUF: &QuickUnionUF{
			parent: make([]int, n),
			count:  n,
		},
	}
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (u *QuickUnionSizeUF) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	// 把size小的根节点链接到size大的根节点，以期望获得更小高度的树
	if u.size[pRoot] < u.size[qRoot] {
		u.parent[pRoot] = qRoot
		u.size[qRoot] += u.size[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.size[pRoot] += u.size[qRoot]
	}
	u.count--
}
