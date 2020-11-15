/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 15:25
 */

package fundamentals

// quick union基于rank的优化, rank即根节点的高度
// 把rank小的根节点链接到rank大的根节点，这比size优化更加合理
type QuickUnionRankUF struct {
	*QuickUnionUF
	rank []int
}

func NewQuickUnionRankUF(n int) *QuickUnionRankUF {
	uf := &QuickUnionRankUF{
		QuickUnionUF: &QuickUnionUF{
			parent: make([]int, n),
			count:  n,
		},
	}
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (u *QuickUnionRankUF) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	// 将rank小的根节点链接到rank大的根节点
	// 这里对辅助数组rank的维护和size不太一样
	// 只有合并后树真的变高了，才需要更新
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		// u.rank[pRoot] == u.rank[qRoot]
		u.parent[pRoot] = qRoot
		u.rank[qRoot]++
	}
	u.count--
}
