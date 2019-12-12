package main

import (
	. "algs4/priorityQueue"
	"fmt"
)

func main() {
	// Insert a bunch of strings
	strings := []string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexPQ(len(strings))
	for idx, s := range strings {
		pq.Insert(idx, s)
	}

	iterator := pq.Yield()
	for hasNext, val := iterator(); hasNext; hasNext, val = iterator() {
		fmt.Println(val, strings[val.(int)])
	}

}
