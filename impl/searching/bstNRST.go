/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/19 11:33
 */

package searching

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

// 迭代版本BST
type BST2 struct {
	root *abstract.TreeNode
}

func NewBST2() *BST2 { return &BST2{} }

func (t *BST2) Put(key, val interface{}) {
	k := abstract.NewTreeNode(key, val)
	if t.root == nil {
		t.root = k
		return
	}

	path := fundamentals.NewLinkedStack()
	node := t.root
	for node != nil {
		path.Push(node)
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			node.Val = val
			return
		} else if cmp < 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	// node == nil
	p := path.Peek().(*abstract.TreeNode)
	if utils.CompareTo(key, p.Key) > 0 {
		p.Right = k
	} else {
		p.Left = k
	}

	for !path.IsEmpty() {
		n := path.Pop().(*abstract.TreeNode)
		n.Size = n.Size + 1
	}
}

func (t *BST2) Get(key interface{}) interface{} {
	node := t.root
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			return node.Val
		} else if cmp < 0 {
			node = node.Left
		} else { //cmp>0
			node = node.Right
		}
	}
	return nil
}

func (t *BST2) Size() int { return t.size(t.root) }

func (t *BST2) size(node *abstract.TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func (t *BST2) Contains(key interface{}) bool { return t.Get(key) != nil }

func (t *BST2) IsEmpty() bool { return t.Size() == 0 }

func (t *BST2) Min() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Min() with empty symbol table")
	return t.min(t.root).Key
}

func (t *BST2) min(node *abstract.TreeNode) *abstract.TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	// node.left == nil
	return node
}

func (t *BST2) Max() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Max() with empty symbol table")
	return t.max(t.root).Key
}

func (t *BST2) max(node *abstract.TreeNode) *abstract.TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	// node.right == nil
	return node
}

// 小于等于key的最大键
func (t *BST2) Floor(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Floor() with empty symbol table")
	x := t.floor(t.root, key)
	utils.AssertF(x != nil, "argument to Floor() is too small")
	return x.Key
}

func (t *BST2) floor(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	var flr *abstract.TreeNode
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			return node
		} else if cmp < 0 {
			node = node.Left
		} else {
			flr = node
			node = node.Right
		}
	}
	return flr
}

func (t *BST2) Ceiling(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Ceiling() with empty symbol table")
	x := t.ceiling(t.root, key)
	utils.AssertF(x != nil, "argument to Ceiling() is too large")
	return x.Key
}

func (t *BST2) ceiling(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	var cil *abstract.TreeNode
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			return node
		} else if cmp > 0 {
			node = node.Right
		} else {
			cil = node
			node = node.Left
		}
	}
	return cil
}

func (t *BST2) Rank(key interface{}) int {
	utils.AssertF(key != nil, "Key is nil")
	return t.rank(t.root, key)
}

func (t *BST2) rank(node *abstract.TreeNode, key interface{}) int {
	b := 0
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			b += t.size(node.Left)
			return b
		} else if cmp < 0 {
			node = node.Left
		} else {
			node = node.Right
			b = b + 1 + t.size(node.Left)
		}
	}
	return b
}

func (t *BST2) Select(k int) interface{} {
	utils.AssertF(k >= 0 && k < t.Size(), "invalid k")
	return t._select(t.root, k)
}

func (t *BST2) _select(node *abstract.TreeNode, k int) interface{} {
	for node != nil {
		leftSize := t.size(node.Left)
		if leftSize == k {
			return node.Key
		} else if leftSize < k {
			node = node.Right
			k = k - leftSize - 1
		} else {
			node = node.Left
		}
	}
	return nil
}

//func (t *BST2) RangeKeys(lo, hi interface{}) abstract.Iterator {
//	queue := fundamentals.NewLinkedQueue()
//	t.rangeKeys(t.root, queue, lo, hi)
//	return queue.Iterate()
//}

//func (t *BST2) rangeKeys(node *abstract.TreeNode, queue abstract.Queue, lo, hi interface{}){
//	stack := fundamentals.NewLinkedStack()
//
//	cmpLo:= utils.CompareTo(lo, node.Key)
//	cmpHi:= utils.CompareTo(hi, node.Key)
//
//
//
//}
