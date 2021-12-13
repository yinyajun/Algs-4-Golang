/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/9/8 14:26
 */

package string

// 三向单词查找树

type Node2 struct {
	c                uint8
	left, mid, right *Node2
	val              interface{}
}

type TST struct {
	root *Node2
}

func (t *TST) get(x *Node2, key string, d int) *Node2 {
	if x == nil {
		return nil
	}

	c := key[d]
	if c < x.c {
		return t.get(x.left, key, d)
	} else if c > x.c {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	} else {
		return x
	}
}
