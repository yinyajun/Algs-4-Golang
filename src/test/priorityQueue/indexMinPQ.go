package main

import (
	. "algs4/priorityQueue"
	"fmt"
)

/**
* go run src/test/indexMinPQ.go
* 3 best
* 0 it
* 6 it
* 4 of
* 8 the
* 2 the
* 5 times
* 7 was
* 1 was
* 9 worst
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	// Insert a bunch of strings
	strings := []string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings))
	for idx, s := range strings {
		pq.Insert(idx, s)
	}
	pq.IncreaseKey(0, "zzz")

	it := pq.Iterate()
	for val := it.Next(); val != nil; val = it.Next() {
		fmt.Println(val, pq.KeyOf(val.(int)))
	}
}
