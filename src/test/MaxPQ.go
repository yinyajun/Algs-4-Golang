package main

import (
	"os"
	"fmt"

	. "algs4/priorityQueue"
	. "util"
)

/**
* $ go run src/test/MaxPQ.go < data/tinyPQ.txt
* Q
* X
* P
* ( 6  left on PQ)
* P
* M
* L
* E
* E
* A
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	pq := NewMaxPQ()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			//k, _ := strconv.Atoi(item) // input should be int
			pq.Insert(item)
		} else if !pq.IsEmpty() {
			fmt.Println(pq.DelMax(), " ")
		}
	}
	fmt.Println("(", pq.Size(), " left on PQ)")
	PrintIterator(pq)
}
