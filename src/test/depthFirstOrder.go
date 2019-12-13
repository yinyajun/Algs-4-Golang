package main

import (
	"algs4/graph"
	"fmt"
	"os"
	"util"
)

/**
* $ go run src/test/depthFirstOrder.go < data/tinyDAG.txt
*    v  pre post
* --------------
*    0    0    8
*    1    3    2
*    2    9   10
*    3   10    9
*    4    2    0
*    5    1    1
*    6    4    7
*    7   11   11
*    8   12   12
*    9    5    6
*   10    8    5
*   11    6    4
*   12    7    3
* Preorder: 0 5 4 1 6 9 11 12 10 2 3 7 8
* Postorder: 4 5 1 12 11 10 9 6 0 3 2 7 8
* Reverse postorder: 8 7 2 3 0 6 9 10 11 12 1 5 4
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := util.NewIn(os.Stdin)
	g := graph.NewDigraphWithIn(in)

	dfs := graph.NewDepthFirstOrder(g)
	fmt.Println("   v  pre post")
	fmt.Println("--------------")
	for v := 0; v < g.V(); v++ {
		fmt.Printf("%4d %4d %4d\n", v, dfs.Pre(v), dfs.Post(v))
	}

	fmt.Print("Preorder: ")
	pre := dfs.PreOrder()
	for ok, v := pre(); ok; ok, v = pre() {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Postorder: ")
	post := dfs.PostOrder()
	for ok, v := post(); ok; ok, v = post() {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Reverse postorder: ")
	reversePost := dfs.ReversePostOrder()
	for ok, v := reversePost(); ok; ok, v = reversePost() {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
