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
	preOrderThread(node.Left, pre)
	preOrderThread(node.Right, pre)
}

func BuildPreOrderThread(root *abstract.TreeNode) *abstract.TreeNode {
	head := abstract.NewTreeNode(nil, nil)
	head.LTag, head.RTag = true, true
	if root == nil {
		head.Left, head.Right = head, head
	}
	head.Left = root
	pre := head
	preOrderThread(root, &pre)
	pre.RTag = true
	pre.Right = head
	head.Right = pre
	return head
}

func PreOrderThreadTraverse(head *abstract.TreeNode) {
	cur := head.Left
	for cur != head {
		Visit(cur)
		for cur.LTag == false {
			cur = cur.Left
			Visit(cur)
		}
		for cur.RTag == true && cur.Right != head {
			cur = cur.Right
		}
		cur = cur.Right
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
	head.Left = root
	inOrderThread(root, &pre)
	pre.RTag = true // 注意此时最后一个节点刚变为pre，并没有线索化右节点
	pre.Right = head
	head.Right = pre
	return head
}

func InOrderThreadTraverse(head *abstract.TreeNode) {
	cur := head.Left
	for cur != head {
		for cur.LTag == false { // 走到左子树尽头，不可能遇到head
			cur = cur.Left
		}
		Visit(cur)
		// 先访问线索构成的后续节点（相当于pop，回到pre）
		for cur.RTag == true && cur.Right != head {
			cur = cur.Right
			Visit(cur)
		}
		// 然后访问正常右子树节点
		cur = cur.Right
	}
}
