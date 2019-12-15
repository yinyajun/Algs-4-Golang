package main

import (
	"fmt"
	"os"

	. "algs4/priorityQueue"
	. "util"
)

/**
* $ go run src/test/MinPQ.go < data/tinyPQ.txt
* E
* A
* E
* ( 6 left on minPQ)
* L
* M
* P
* P
* Q
* X
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	pq := NewMinPQ()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			pq.Insert(item)
		} else if !pq.IsEmpty() {
			fmt.Println(pq.DelMin(), " ")
		}
	}
	fmt.Println("(", pq.Size(), "left on minPQ)")
	PrintIterator(pq.Iterate())
}
