package main

import (
	. "algs4/priorityQueue"
	"fmt"
)

type Iterator interface {
	next() interface{}
}

func next(pq *IndexPQ) (interface{}, bool) {
	if !pq.IsEmpty() {
		return pq.DelMax(), !pq.IsEmpty()
	}
	return nil, false
}

func main() {

	strings := []string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexPQ(len(strings))
	for idx, s := range strings {
		pq.Insert(idx, s)
	}

	for val, hasNext := next(pq); hasNext; val, hasNext = next(pq) {
		fmt.Println(val)
	}

}
