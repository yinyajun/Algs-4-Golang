package main

import (
	"fmt"

	. "algs4/priorityQueue"
)

func main() {
	// Insert a bunch of strings
	strings := []string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexPQ(len(strings))
	for idx, s := range strings {
		pq.Insert(idx, s)
	}

	// print each key using the Iterator
	for _, i := range pq.Iterator() {
		fmt.Println(i, strings[i])
	}
}
