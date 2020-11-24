/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/24 11:33
 */

package main

import (
	"Algs-4-Golang/impl/sorting"
	"Algs-4-Golang/utils"

	"math/rand"
	"time"
)

type SortCompare struct{}

func (c *SortCompare) time(alg string, a []int) time.Duration {
	start := time.Now()
	switch alg {
	case "Insertion":
		sorting.NewInsertion().Sort(a, func(i, j int) bool { return a[i] < a[j] })
	case "Selection":
		sorting.NewSelection().Sort(a, func(i, j int) bool { return a[i] < a[j] })
	case "AdvancedInsertion":
		sorting.NewAdvancedInsertionSorter().Sort(a, func(i, j int) bool { return a[i] < a[j] })
	case "Shell":
		sorting.NewShell().Sort(a, func(i, j int) bool { return a[i] < a[j] })
	}
	return time.Since(start)
}

func (c *SortCompare) timeRandomInput(alg string, N int, T int) time.Duration {
	var total time.Duration
	a := make([]int, N)
	for t := 0; t < T; t++ {
		for i := 0; i < N; i++ {
			a[i] = rand.Intn(1e8)
		}
		total += c.time(alg, a)
	}
	return total
}

func main() {
	t := new(SortCompare)
	N := 10000
	T := 10
	t1 := t.timeRandomInput("Insertion", N, T)
	t2 := t.timeRandomInput("Selection", N, T)
	t3 := t.timeRandomInput("Shell", N, T)
	utils.StdOut.Println(t1, t2, t3)
}
