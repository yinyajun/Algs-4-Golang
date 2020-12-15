/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/12/12 11:36
 */

package searching

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/utils"
)

type Node struct {
	key         interface{}
	val         interface{}
	left, right abstract.Node2
	size        int
}

func NewNode(key, val interface{}, left, right abstract.Node2, size int) *Node {
	return &Node{key, val, left, right, size}
}

func (n *Node) Key() interface{}      { return n.key }
func (n *Node) Value() interface{}    { return n.val }
func (n *Node) Left() abstract.Node2  { return n.left }
func (n *Node) Right() abstract.Node2 { return n.right }
func (n *Node) Size() int             { return n.size }

func (n *Node) SetKey(key interface{})       { n.key = key }
func (n *Node) SetValue(val interface{})     { n.val = val }
func (n *Node) SetLeft(node abstract.Node2)  { n.left = node }
func (n *Node) SetRight(node abstract.Node2) { n.right = node }
func (n *Node) SetSize(size int)             { n.size = size }

type BST struct {
	root abstract.Node2
}

func NewBST() *BST { return &BST{} }

func (t *BST) Put(key, val interface{}) {
	utils.AssertF(key != nil, "Key is nil")

	if val == nil {
		t.Delete(key)
	}

	t.root = t.put(t.root, key, val)
}

// 在以node为根的子树中添加键值
func (t *BST) put(node abstract.Node2, key, val interface{}) abstract.Node2 {
	if node == nil {
		return NewNode(key, val, nil, nil, 1)
	}

	cmp := utils.CompareTo(node.Key(), key)

	if cmp > 0 {
		node.SetLeft(t.put(node.Left(), key, val))
	} else if cmp < 0 {
		node.SetRight(t.put(node.Right(), key, val))
	} else {
		node.SetValue(val)
	}
	node.SetSize(1 + t.size(node.Left()) + t.size(node.Right()))
	return node
}

func (t *BST) Get(key interface{}) interface{} {
	return t.get(t.root, key)
}

func (t *BST) get(node abstract.Node2, key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")

	if node == nil {
		return nil
	}

	cmp := utils.CompareTo(node.Key(), key)
	if cmp < 0 {
		return t.get(node.Right(), key)
	} else if cmp > 0 {
		return t.get(node.Left(), key)
	} else {
		return node.Value()
	}
}

func (t *BST) Delete(key interface{}) {}

func (t *BST) delete(node abstract.Node2, key interface{}) {
	if node == nil {
		return
	}
	cmp := utils.CompareTo(node.Value(), key)
	if cmp < 0 {
	} else if cmp > 0 {

	} else {

	}
}

func (t *BST) Contains(key interface{}) bool {
	utils.AssertF(key != nil, "Key is nil")
	return t.Get(key) != nil
}

func (t *BST) IsEmpty() bool { return t.Size() == 0 }

func (t *BST) Size() int { return t.size(t.root) }

func (t *BST) size(node abstract.Node2) int {
	if node == nil {
		return 0
	}
	return node.Size()
}

func (t *BST) Min() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Min() with empty symbol table")
	return t.min(t.root).Key()
}

func (t *BST) min(node abstract.Node2) abstract.Node2 {
	if node.Left() == nil {
		return node
	}
	return t.min(node.Left())
}

// 最大的键
func (t *BST) Max() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Max() with empty symbol table")
	return t.max(t.root).Key()
}

// 如果右子树为空，那么最大键就是根节点
// 右子树不为空，最大键就是右子树中的最大键
func (t *BST) max(node abstract.Node2) abstract.Node2 {
	if node.Right() == nil {
		return node
	}
	return t.max(node.Right())
}

// 小于等于key的最大键
func (t *BST) Floor(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")

}

// 如果key小于node的key，floor值一定在左子树中
// 如果key大于node的key，如果右子树中存在小于等于key的节点时，floor值在右子树中，否则floor值为根节点
func (t *BST) floor(node abstract.Node2, key interface{}) abstract.Node2 {
	if node == nil {
		return nil
	}
	cmp := utils.CompareTo(node.Key(), key)
	if cmp == 0 {
		return node
	} else if cmp > 0 {
		return t.floor(node.Left(), key)
	}
	// key > node.Key()
	if n := t.floor(node.Right(), key); n != nil {
		return n
	}
	return node
}

// 大于等于key的最小键
func (t *BST) Ceiling(key interface{}) interface{} {}

func (t *BST) Rank(key interface{}) int {}

func (t *BST) Select(k int) interface{} {}

func (t *BST) DeleteMin() {}

func (t *BST) DeleteMax() {}

func (t *BST) RangeSize(lo, hi interface{}) int {}

func (t *BST) RangeKeys(lo, hi interface{}) abstract.Iterator {}

func (t *BST) Keys() abstract.Iterator {}
