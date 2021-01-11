/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:54
 */

package abstract

type Node struct {
	Key, Val interface{}
	Next     *Node
}

type TreeNode struct {
	Key, Val    interface{}
	Left, Right *TreeNode
	Size        int
	LTag, RTag  bool
}

func NewTreeNode(key, val interface{}) *TreeNode {
	return &TreeNode{Key: key, Val: val, Size: 1}
}
