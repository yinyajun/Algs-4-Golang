/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/1/8 17:04
 */

package searching

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
)

var Result strings.Builder

// -------------------
// PreOrder
// -------------------

func PreOrder(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	Visit(root)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PreOrderNR(root *abstract.TreeNode) {
	s := fundamentals.NewLinkedStack()
	cur := root

	for cur != nil || !s.IsEmpty() {
		if cur != nil { // 访问当前合法节点
			Visit(cur)
			s.Push(cur)
			cur = cur.Left
		} else { // cur==nil，此时需要通过pop回溯到上一个节点
			cur = s.Pop().(*abstract.TreeNode)
			cur = cur.Right
		}
	}
}

func PreOrderNR2(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	s := fundamentals.NewLinkedStack()
	s.Push(root)
	for !s.IsEmpty() { // 根-右-左 => 根-左-右
		cur := s.Pop().(*abstract.TreeNode)
		Visit(cur)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

// 一路向左的过程中建立线索，然后回溯的时候取消线索
// 需要证明的是，当cur.left == nil的时候，cur.right一定不为空
func PreOrderMorris(root *abstract.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			Visit(cur)
			cur = cur.Right // cur的right一定有值（线索or link）
		} else {
			// cur.left != nil
			// 找到中序的前驱节点
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			// 构建线索
			if pre.Right == nil {
				Visit(cur)
				pre.Right = cur
				cur = cur.Left
			} else {
				// 取消线索
				pre.Right = nil
				cur = cur.Right
			}
		}
	}
}

func PreOrderThread(root *abstract.TreeNode) {
	head := BuildPreOrderThread(root)
	PreOrderThreadTraverse(head)
}

// -------------------
// InOrder
// -------------------

func InOrder(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	Visit(root)
	InOrder(root.Right)
}

func InOrderNR(root *abstract.TreeNode) {
	cur := root
	s := fundamentals.NewLinkedStack()
	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else { // cur == nil
			cur = s.Pop().(*abstract.TreeNode)
			Visit(cur)
			cur = cur.Right
		}
	}
}

// 一路向左的过程中建立线索，然后回溯的时候取消线索
func InOrderMorris(root *abstract.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			Visit(cur)
			cur = cur.Right
		} else { // cur.Left != nil
			// 找到当前节点的中序前驱节点（左子树的最右边节点）
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur { // 右子树不为空，且不能是线索
				pre = pre.Right
			}
			// 将当前节点作为前驱节点的右孩子（线索）
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else { // 将线索取消
				pre.Right = nil
				Visit(cur)
				cur = cur.Right
			}
		}
	}
}

func InOrderThread(root *abstract.TreeNode) {
	root = BuildInOrderThread(root)
	InOrderThreadTraverse(root)
}

// -------------------
// PostOrder
// -------------------

func PostOrder(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	Visit(root)
}

func PostOrderNR(root *abstract.TreeNode) {
	cur := root
	var pre *abstract.TreeNode
	s := fundamentals.NewLinkedStack()

	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else { // cur == nil
			cur = s.Peek().(*abstract.TreeNode)       // 此时不能急于pop出该节点，由于post的特殊，遍历完右子树仍需要返回该节点
			if cur.Right == nil || cur.Right == pre { // 当前节点右子树为空或者右子树已经遍历过，此时相当于右子树不需要遍历
				Visit(cur)
				s.Pop() //  弹出cur节点
				pre = cur
				cur = nil // 当前节点生命周期结束，需要继续pop
			} else { //右子树需要遍历
				cur = cur.Right
			}
		}
	}
}

func PostOrderNR_(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	s := fundamentals.NewLinkedStack()
	var cur *abstract.TreeNode
	pre := root
	s.Push(root)
	for !s.IsEmpty() {
		cur = s.Peek().(*abstract.TreeNode)
		if cur.Left != nil && cur.Left != pre && cur.Right != pre { // 左子树不为空，且左右子树均未遍历过
			s.Push(cur.Left)
		} else if cur.Right != nil && cur.Right != pre { // 右子树为空，且右子树未遍历过
			s.Push(cur.Right)
		} else {
			Visit(s.Pop().(*abstract.TreeNode))
			pre = cur
		}
	}
}

func PostOrderNR2(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	s := fundamentals.NewLinkedStack()
	ret := fundamentals.NewLinkedStack()
	s.Push(root)
	for !s.IsEmpty() { // 根-左-右 =》 根-右-左 =》 左-右-根
		cur := s.Pop().(*abstract.TreeNode)
		ret.Push(cur)
		if cur.Left != nil {
			s.Push(cur.Left)
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
	}
	iter := ret.Iterate()
	for iter.First(); iter.HasNext(); {
		Visit(iter.Next().(*abstract.TreeNode))
	}
}

// 依次逆序打印左子树的右边界
func PostOrderMorris(root *abstract.TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
		} else { // cur.Left != nil
			// 找到当前节点的中序前驱节点（左子树的最右边节点）
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur { // 右子树不为空，且不能是线索
				pre = pre.Right
			}
			// 将当前节点作为前驱节点的右孩子（线索）
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else { // 将线索取消
				pre.Right = nil
				printNode(cur.Left) // 发生的时机是在第二次回溯到cur，而不是在第三次回溯。那么此时只能将左子树逆序打印
				cur = cur.Right
			}
		}
	}
	printNode(root)
}

// 类似非递归反转链表，不过这里是右子树这条线（相当于链表）
func reverseNode(node *abstract.TreeNode) *abstract.TreeNode {
	var pre, next *abstract.TreeNode
	for node != nil {
		next = node.Right
		node.Right = pre
		pre, node = node, next
	}
	// node== nil
	return pre
}

// 逆序打印，反转-打印-再反转
func printNode(node *abstract.TreeNode) {
	tail := reverseNode(node)
	cur := tail
	for cur != nil {
		Visit(cur)
		cur = cur.Right
	}
	reverseNode(tail)
}

// 这种方式很讨巧，但是使morris遍历丧失了其o(1)空间复杂度的优势
func PostOrderMorris2(root *abstract.TreeNode) {
	// 左右根=》根右左
	res := fundamentals.NewLinkedStack()
	cur := root
	for cur != nil {
		if cur.Right == nil {
			res.Push(cur)
			cur = cur.Left
		} else { //cur.right != nil
			pre := cur.Right
			for pre.Left != nil && pre.Left != cur {
				pre = pre.Left
			}
			if pre.Left == nil {
				res.Push(cur)
				pre.Left = cur
				cur = cur.Right
			} else { // pre.left == cur
				pre.Left = nil
				cur = cur.Left
			}
		}
	}
	it := res.Iterate()
	for it.First(); it.HasNext(); {
		Visit(it.Next().(*abstract.TreeNode))
	}
}

// util
func GetFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	names := strings.Split(name, "/")
	name = names[len(names)-1]
	return fmt.Sprintf("%-30s| ", name)
}

func (r *Results) Add(root *abstract.TreeNode, f func(node *abstract.TreeNode)) {
	Result = strings.Builder{}
	Result.WriteString(GetFunctionName(f))
	f(root)
	*r = append(*r, Result.String())
}

type Results []string

func Visit(node *abstract.TreeNode) { Result.WriteString(node.Key.(string) + " ") }
