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
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

// 递归版本BST
type BST struct {
	root *abstract.TreeNode
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
func (t *BST) put(node *abstract.TreeNode, key, val interface{}) *abstract.TreeNode {
	if node == nil {
		return abstract.NewTreeNode(key, val)
	}

	cmp := utils.CompareTo(node.Key, key)

	if cmp > 0 {
		node.Left = t.put(node.Left, key, val)
	} else if cmp < 0 {
		node.Right = t.put(node.Right, key, val)
	} else {
		node.Val = val
	}
	node.Size = 1 + t.size(node.Left) + t.size(node.Right)
	return node
}

func (t *BST) Get(key interface{}) interface{} {
	return t.get(t.root, key)
}

func (t *BST) get(node *abstract.TreeNode, key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")

	if node == nil {
		return nil
	}

	cmp := utils.CompareTo(node.Key, key)
	if cmp < 0 {
		return t.get(node.Right, key)
	} else if cmp > 0 {
		return t.get(node.Left, key)
	} else {
		return node.Val
	}
}

func (t *BST) Contains(key interface{}) bool {
	utils.AssertF(key != nil, "Key is nil")
	return t.Get(key) != nil
}

func (t *BST) IsEmpty() bool { return t.Size() == 0 }

func (t *BST) Size() int { return t.size(t.root) }

func (t *BST) size(node *abstract.TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

// 最小的键
func (t *BST) Min() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Min() with empty symbol table")
	return t.min(t.root).Key
}

// 如果左子树为空，那么最小健就是根节点
// 左子树不为空，最小键就是左子树中的最小键
func (t *BST) min(node *abstract.TreeNode) *abstract.TreeNode {
	if node.Left == nil {
		return node
	}
	return t.min(node.Left)
}

// 最大的键
func (t *BST) Max() interface{} {
	utils.AssertF(!t.IsEmpty(), "called Max() with empty symbol table")
	return t.max(t.root).Key
}

// 如果右子树为空，那么最大键就是根节点
// 右子树不为空，最大键就是右子树中的最大键
func (t *BST) max(node *abstract.TreeNode) *abstract.TreeNode {
	if node.Right == nil {
		return node
	}
	return t.max(node.Right)
}

// 小于等于key的最大键,<-|
func (t *BST) Floor(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Floor() with empty symbol table")
	x := t.floor(t.root, key)
	utils.AssertF(x != nil, "argument to Floor() is too small")
	return x.Key
}

// 如果key小于node的key，floor值一定在左子树中
// 如果key等于node的key，就是floor值
// 如果key大于node的key，如果右子树中存在小于等于key的节点时，floor值在右子树中，否则floor值为根节点
func (t *BST) floor(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	if node == nil {
		return nil
	}
	cmp := utils.CompareTo(node.Key, key)
	if cmp == 0 {
		return node
	} else if cmp > 0 {
		return t.floor(node.Left, key)
	}
	// key > node.Key()
	if n := t.floor(node.Right, key); n != nil {
		return n
	}
	return node
}

// 大于等于key的最小键，|->
func (t *BST) Ceiling(key interface{}) interface{} {
	utils.AssertF(key != nil, "Key is nil")
	utils.AssertF(!t.IsEmpty(), "calls Ceiling() with empty symbol table")
	x := t.ceiling(t.root, key)
	utils.AssertF(x != nil, "argument to Ceiling() is too large")
	return x.Key
}

// 如果key大于node的key，ceiling值一定在右子树中
// 如果key等于node的key，就是ceiling值
// 如果key小于node的key，如果左子树中存在大于等于key的节点时，ceiling值在左子树中，否则floor值为根节点
func (t *BST) ceiling(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	if node == nil {
		return nil
	}
	cmp := utils.CompareTo(key, node.Key)
	if cmp == 0 {
		return node
	} else if cmp > 0 {
		return t.ceiling(node.Right, key)
	} else { // cmp<0
		if n := t.ceiling(node.Left, key); n != nil {
			return n
		}
		return node
	}
}

func (t *BST) Rank(key interface{}) int {
	utils.AssertF(key != nil, "Key is nil")
	return t.rank(t.root, key)
}

func (t *BST) rank(node *abstract.TreeNode, key interface{}) int {
	if node == nil {
		return 0
	}
	cmp := utils.CompareTo(key, node.Key)
	if cmp < 0 {
		return t.rank(node.Left, key)
	} else if cmp > 0 {
		return t.size(node.Left) + 1 + t.rank(node.Right, key)
	} else {
		return t.size(node.Left)
	}
}

func (t *BST) Select(k int) interface{} {
	utils.AssertF(k >= 0 && k < t.Size(), "invalid k")
	return t._select(t.root, k).Key
}

func (t *BST) _select(node *abstract.TreeNode, k int) *abstract.TreeNode {
	if node == nil {
		return nil
	}
	leftSize := t.size(node.Left)
	if leftSize > k {
		return t._select(node.Left, k)
	} else if leftSize < k {
		return t._select(node.Right, k-leftSize-1)
	} else {
		return node
	}
}

func (t *BST) DeleteMin() {
	utils.AssertF(!t.IsEmpty(), "called DeleteMin() with empty symbol table")
	t.deleteMin(t.root)
}

func (t *BST) deleteMin(node *abstract.TreeNode) *abstract.TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = t.deleteMin(node.Left)
	node.Size = t.size(node.Left) + t.size(node.Right) + 1
	return node
}

func (t *BST) DeleteMax() {
	utils.AssertF(!t.IsEmpty(), "called DeleteMax() with empty symbol table")
	t.deleteMax(t.root)
}

func (t *BST) deleteMax(node *abstract.TreeNode) *abstract.TreeNode {
	if node.Right == nil {
		return node.Left
	}
	node.Right = t.deleteMax(node.Right)
	node.Size = t.size(node.Left) + t.size(node.Right) + 1
	return node
}

func (t *BST) Delete(key interface{}) {
	t.delete(t.root, key)
}

func (t *BST) delete(node *abstract.TreeNode, key interface{}) *abstract.TreeNode {
	if node == nil {
		return nil
	}
	cmp := utils.CompareTo(key, node.Key)
	if cmp < 0 {
		node.Left = t.delete(node.Left, key)
	} else if cmp > 0 {
		node.Right = t.delete(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// 用右子树中最小节点代替当前节点，右子树设为删除过最小节点的右子树，左子树设为原先节点的左子树
		d := node
		node = t.min(d.Right)
		node.Right = t.deleteMin(d.Right)
		node.Left = d.Left
	}
	node.Size = t.size(node.Left) + t.size(node.Right) + 1
	return node
}
func (t *BST) RangeSize(lo, hi interface{}) int {
	return t.rangeSize(t.root, lo, hi)
}

func (t *BST) rangeSize(node *abstract.TreeNode, lo, hi interface{}) int {
	if node == nil {
		return 0
	}
	cmpLo := utils.CompareTo(lo, node.Key)
	cmpHi := utils.CompareTo(hi, node.Key)

	size := 0
	if cmpLo < 0 {
		size += t.rangeSize(node.Left, lo, hi)
	}

	if cmpLo <= 0 && cmpHi >= 0 {
		size += 1
	}
	if cmpHi > 0 {
		size += t.rangeSize(node.Right, lo, hi)
	}
	return size
}

func (t *BST) RangeKeys(lo, hi interface{}) abstract.Iterator {
	queue := fundamentals.NewLinkedQueue()
	t.rangeKeys(t.root, queue, lo, hi)
	return queue.Iterate()
}

func (t *BST) rangeKeys(node *abstract.TreeNode, queue abstract.Queue, lo, hi interface{}) {
	if node == nil {
		return
	}
	cmpLo := utils.CompareTo(lo, node.Key)
	cmpHi := utils.CompareTo(hi, node.Key)
	if cmpLo < 0 {
		t.rangeKeys(node.Left, queue, lo, hi)
	}

	if cmpLo <= 0 && cmpHi >= 0 {
		queue.Enqueue(node.Key)
	}
	if cmpHi > 0 {
		t.rangeKeys(node.Right, queue, lo, hi)
	}

}

func (t *BST) Keys() abstract.Iterator {
	return t.RangeKeys(t.Min(), t.Max())
}
