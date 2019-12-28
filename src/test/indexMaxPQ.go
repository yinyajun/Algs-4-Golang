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
	pq := NewIndexMaxPQ(len(strings))
	for idx, s := range strings {
		pq.Insert(idx, s)
	}

	pq.IncreaseKey(0, "zz")

	// print each key using the iterator
	it := pq.Iterate()
	for val := it.Next(); val != nil; val = it.Next() {
		fmt.Println(val, pq.KeyOf(val.(int)))
	}

}
