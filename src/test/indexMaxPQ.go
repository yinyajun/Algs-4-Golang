package main

import (
	. "algs4/priorityQueue"
	"fmt"
)

/**
* $ go run src/test/indexMaxPQ.go
* 9 worst
* 1 was
* 7 was
* 5 times
* 8 the
* 2 the
* 4 of
* 6 it
* 0 it
* 3 best
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/


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
