package main

import (
	"fmt"
	"os"

	. "algs4/util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type WeightedQuickUnionUF struct {
	sz  []int // 各个根节点所对应的分量的大小
	id  []int // 父链接数组
	cnt int
}

func NewWeightedQuickUnionUF(N int) *WeightedQuickUnionUF {
	id := make([]int, N)
	sz := make([]int, N)
	for idx := range id {
		id[idx] = idx
	}
	for idx := range sz {
		sz[idx] = 1
	}
	return &WeightedQuickUnionUF{
		id:  id,
		cnt: N,
		sz:  sz,
	}
}

func (m *WeightedQuickUnionUF) count() int {
	return m.cnt
}

func (m *WeightedQuickUnionUF) connected(p int, q int) bool {
	return m.find(p) == m.find(q)
}

func (m *WeightedQuickUnionUF) find(p int) int {
	for p != m.id[p] {
		p = m.id[p]
	}
	return p
}

func (m *WeightedQuickUnionUF) union(p int, q int) {
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

func main() {
	in := NewIn(os.Stdin)
	N := in.ReadInt()
	uf := NewWeightedQuickUnionUF(N)
	for in.HasNext() {
		p := in.ReadInt()
		q := in.ReadInt()
		if uf.connected(p, q) {
			continue
		}
		uf.union(p, q)
		//fmt.Println(p, q)
	}
	fmt.Println(uf.count(), "components")
}
