/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/1/7 22:10
 */

package searching

import (
	"Algs-4-Golang/abstract"
)

// -------------------
// PreOrder
// -------------------
func preOrderThread(node *abstract.TreeNode, pre **abstract.TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil {
		node.LTag = true
		node.Left = *pre
	}
	if (*pre).Right == nil {
		(*pre).RTag = true
		(*pre).Right = node
	}
	*pre = node

	if node.LTag == false { // 这里和中序线索化不同，因为left可能是刚建好的线索
		preOrderThread(node.Left, pre)
	}
	if node.RTag == false { //这里和中序线索化不同，因为right可能是刚建好的线索
		preOrderThread(node.Right, pre)
	}
}

func BuildPreOrderThread(root *abstract.TreeNode) *abstract.TreeNode {
	head := abstract.NewTreeNode(nil, nil) // 虚拟头结点
	head.LTag, head.RTag = true, true
	if root == nil {
		head.Left, head.Right = head, head
	}
	pre := head
	preOrderThread(root, &pre) // 此时head的left，pre的right没有设置
	head.Left, pre.Right = pre, head
	return head
}

// 前序遍历线索二叉树（根-左-右）
// 1. 访问根节点
// 2. 如果左子树非空，进入左子树的根节点
// 3. 如果右子树非空，进入右子树的根节点
// 4. 如果右子树为空（线索），进入右thread提供的后续节点（相当于回溯）

func PreOrderThreadTraverse(head *abstract.TreeNode) {
	cur := head.Right
	for cur != head {
		Visit(cur)
		if cur.LTag == false {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
}

// -------------------
// InOrder
// -------------------
func inOrderThread(node *abstract.TreeNode, pre **abstract.TreeNode) {
	if node == nil {
		return
	}
	inOrderThread(node.Left, pre)
	if node.Left == nil {
		node.LTag = true
		node.Left = *pre
	}
	if (*pre).Right == nil {
		(*pre).RTag = true
		(*pre).Right = node
	}
	*pre = node // pre记录上次遍历的非空节点
	inOrderThread(node.Right, pre)
}

func BuildInOrderThread(root *abstract.TreeNode) *abstract.TreeNode {
	head := abstract.NewTreeNode(nil, nil)
	head.LTag, head.RTag = true, true
	if root == nil {
		head.Left, head.Right = head, head
		return head
	}
	pre := head
	inOrderThread(root, &pre) // 此时head的left，pre的right没有设置
	head.Left, pre.Right = pre, head
	return head
}

// 中序遍历线索二叉树（左根右）退化为： 左（根） -右
// 1. 初始化为最左边节点为根节点
// 2. 访问根节点
// 3. 如果右子树不为空，进入右子树的根节点
// 4. 如果右子树为空（线索），进入右thread提供的后续节点（相当于回溯）

func InOrderThreadTraverse(head *abstract.TreeNode) {
	cur := head.Right
	for cur != head {
		Visit(cur)
		cur = cur.Right
	}
}
