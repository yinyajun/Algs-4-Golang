package main

import (
	"os"
	"strconv"
	"fmt"

	. "algs4/priorityQueue"
	. "util"
)

func main() {
	pq := NewMaxPQ()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			k, _ := strconv.Atoi(item) // input should be int
			pq.Insert(k)
		} else if !pq.IsEmpty() {
			fmt.Println(pq.DelMax(), " ")
		}
	}
	fmt.Println("(", pq.Size(), " left on PQ)")
	PrintIterator(pq)
}
