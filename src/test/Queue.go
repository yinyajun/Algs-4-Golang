package main

import (
	"os"
	"fmt"

	. "algs4/queue"
	. "util"
)

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
