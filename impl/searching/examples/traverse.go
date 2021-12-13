/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/1/2 15:06
 */

package main

// go run impl/searching/examples/traverse.go
// +----------------------------------------------+
// | PreOrder                                     |
// +----------------------------------------------+
// | searching.PreOrder            | 1 2 4 5 3 7  |
// | searching.PreOrderNR          | 1 2 4 5 3 7  |
// | searching.PreOrderNR2         | 1 2 4 5 3 7  |
// | searching.PreOrderMorris      | 1 2 4 5 3 7  |
// | searching.PreOrderThread      | 1 2 4 5 3 7  |
// +----------------------------------------------+
// +----------------------------------------------+
// | InOrder                                      |
// +----------------------------------------------+
// | searching.InOrder             | 4 2 5 1 3 7  |
// | searching.InOrderNR           | 4 2 5 1 3 7  |
// | searching.InOrderMorris       | 4 2 5 1 3 7  |
// | searching.InOrderThread       | 4 2 5 1 3 7  |
// +----------------------------------------------+
// +----------------------------------------------+
// | PostOrder                                    |
// +----------------------------------------------+
// | searching.PostOrder           | 4 5 2 7 3 1  |
// | searching.PostOrderNR         | 4 5 2 7 3 1  |
// | searching.PostOrderNR_        | 4 5 2 7 3 1  |
// | searching.PostOrderNR2        | 4 5 2 7 3 1  |
// | searching.PostOrderMorris     | 4 5 2 7 3 1  |
// | searching.PostOrderMorris2    | 4 5 2 7 3 1  |
// +----------------------------------------------+
//

import (
	"Algs-4-Golang/impl/searching"
	"Algs-4-Golang/utils"
)

func main() {
	array := []string{"1", "2", "4", "#", "#", "5", "#", "#", "3", "#", "7", "#", "#"}
	tree := searching.CreateTreeFromArray(array)

	// preOrder
	rets := new(searching.Results)
	rets.Add(tree, searching.PreOrder)
	rets.Add(tree, searching.PreOrderNR)
	rets.Add(tree, searching.PreOrderNR2)
	rets.Add(tree, searching.PreOrderMorris)
	rets.Add(searching.CloneTree(tree), searching.PreOrderThread)
	utils.PrintInTable([]string{"PreOrder"}, *rets)

	// inOrder
	rets = new(searching.Results)
	rets.Add(tree, searching.InOrder)
	rets.Add(tree, searching.InOrderNR)
	rets.Add(tree, searching.InOrderMorris)
	rets.Add(searching.CloneTree(tree), searching.InOrderThread)
	utils.PrintInTable([]string{"InOrder"}, *rets)

	// postOrder
	rets = new(searching.Results)
	rets.Add(tree, searching.PostOrder)
	rets.Add(tree, searching.PostOrderNR)
	rets.Add(tree, searching.PostOrderNR_)
	rets.Add(tree, searching.PostOrderNR2)
	rets.Add(tree, searching.PostOrderMorris)
	rets.Add(tree, searching.PostOrderMorris2)
	utils.PrintInTable([]string{"PostOrder"}, *rets)
}
