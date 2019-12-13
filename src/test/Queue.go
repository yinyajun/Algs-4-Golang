package main

import (
	"os"
	"fmt"

	. "algs4/queue"
	. "util"
)

/**
* $ go run  src/test/Queue.go < data/tobe.txt
* to be or not to be ( 2 left on queue)
* that
* is
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

func main() {
	q := NewQueue()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			q.Enqueue(item)
		} else if !q.IsEmpty() {
			fmt.Print(q.Dequeue(), " ")
		}
	}
	fmt.Println("(", q.Size(), "left on queue)")
	PrintIterator(q)
}
