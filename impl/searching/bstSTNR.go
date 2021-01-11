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
type BSTNR struct {
	root *abstract.TreeNode
}

func NewBSTNR() *BSTNR { return &BSTNR{} }

func (t *BSTNR) Put(key, val interface{}) {
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

func (t *BSTNR) Get(key interface{}) interface{} {
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

func (t *BSTNR) Size() int { return t.size(t.root) }

func (t *BSTNR) size(node *abstract.TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func (t *BSTNR) Contains(key interface{}) bool { return t.Get(key) != nil }

func (t *BSTNR) IsEmpty() bool { return t.Size() == 0 }

func (t *BSTNR) Min() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Min() with empty symbol table")
	return t.min(t.root).Key
}

func (t *BSTNR) min(node *abstract.TreeNode) *abstract.TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	// node.left == nil
	return node
}

func (t *BSTNR) Max() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Max() with empty symbol table")
	return t.max(t.root).Key
}

func (t *BSTNR) max(node *abstract.TreeNode) *abstract.TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	// node.right == nil
	return node
}

// 小于等于key的最大键
func (t *BSTNR) Floor(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Floor() with empty symbol table")
	x := t.floor(t.root, key)
	utils.AssertF(x != nil, "argument to Floor() is too small")
	return x.Key
}

func (t *BSTNR) floor(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	var flr *abstract.TreeNode
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp == 0 {
			return node
		} else if cmp < 0 {
			node = node.Left
		} else { // cmp > 0
			flr = node
			node = node.Right
		}
	}
	return flr
}

func (t *BSTNR) Ceiling(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Ceiling() with empty symbol table")
	x := t.ceiling(t.root, key)
	utils.AssertF(x != nil, "argument to Ceiling() is too large")
	return x.Key
}

func (t *BSTNR) ceiling(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
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

func (t *BSTNR) Rank(key interface{}) int {
	utils.AssertF(key != nil, "Key is nil")
	return t.rank(t.root, key)
}

func (t *BSTNR) rank(node *abstract.TreeNode, key interface{}) int {
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

func (t *BSTNR) Select(k int) interface{} {
	utils.AssertF(k >= 0 && k < t.Size(), "invalid k")
	return t._select(t.root, k)
}

func (t *BSTNR) _select(node *abstract.TreeNode, k int) interface{} {
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

func (t *BSTNR) RangeKeys(lo, hi interface{}) abstract.Iterator {
	queue := fundamentals.NewLinkedQueue()
	t.rangeKeys(t.root, queue, lo, hi)
	return queue.Iterate()
}

func (t *BSTNR) rangeKeys(node *abstract.TreeNode, queue abstract.Queue, lo, hi interface{}) {
	stack := fundamentals.NewLinkedStack()

	for node != nil || !stack.IsEmpty() {
		if node != nil {
			stack.Push(node)
			cmpLo := utils.CompareTo(lo, node.Key)
			if cmpLo < 0 {
				node = node.Left
			} else {
				node = nil
			}
		} else { // node == nil
			node = stack.Pop().(*abstract.TreeNode)

			cmpLo := utils.CompareTo(lo, node.Key)
			cmpHi := utils.CompareTo(hi, node.Key)
			// visit node
			if cmpLo <= 0 && cmpHi >= 0 {
				queue.Enqueue(node.Key)
			}
			if cmpHi > 0 {
				node = node.Right
			} else {
				node = nil
			}
		}
	}
}

func (t *BSTNR) Delete(key interface{}) {
	t.delete(&t.root, key)
}

// todo: check
func (t *BSTNR) delete(root **abstract.TreeNode, key interface{}) {
	node := *root
	var parent *abstract.TreeNode
	for node != nil {
		cmp := utils.CompareTo(key, node.Key)
		if cmp < 0 {
			parent = node
			node = node.Left
		} else if cmp > 0 {
			parent = node
			node = node.Right
		} else { // 当前节点为待删除节点 node.key == key
			// 待删除节点无左孩子，直接用右孩子代替
			if node.Left == nil {
				if parent == nil { // node == *root
					*root = node.Right
				} else {
					if node == parent.Left {
						parent.Left = node.Right
					} else {
						parent.Right = node.Right
					}
				}
			} else if node.Left == nil { // 待删除节点无右孩子，直接用左孩子代替
				if parent == nil { // node == *root
					*root = node.Left
				} else {
					if node == parent.Left {
						parent.Left = node.Left
					} else {
						parent.Right = node.Left
					}
				}
			} else { // 待删除节点左右孩子都存在
				min := node.Right
				minPre := node
				for min.Left != nil {
					minPre = min
					min = min.Left
				}
				node.Key = min.Key
				cmp := utils.CompareTo(min.Key, minPre.Key)
				if cmp < 0 {
					minPre.Left = min.Right
				} else {
					minPre.Right = min.Right
				}
			}
		}
	}
}

func (t *BSTNR) Keys() abstract.Iterator {
	return t.RangeKeys(t.Min(), t.Max())
}
