/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/24 11:33
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/sorting"
	"Algs-4-Golang/utils"
	"math/rand"
	"time"
)

type SortCompare struct{}

func (c *SortCompare) time(alg string, a []float64) time.Duration {
	var sorter abstract.Sorter
	start := time.Now()
	switch alg {
	case "Insertion":
		sorter = sorting.NewInsertion()
	case "Selection":
		sorter = sorting.NewSelection()
	case "AdvancedInsertion":
		sorter = sorting.NewAdvancedInsertion()
	case "Shell":
		sorter = sorting.NewShell()
	case "Merge":
		sorter = sorting.NewMerge()
	case "MergeBU":
		sorter = sorting.NewMergeBU()
	case "AdvancedMerge":
		sorter = sorting.NewAdvancedMerge()
	case "Quick":
		sorter = sorting.NewQuick()
	case "ThreeWayQuick":
		sorter = sorting.NewThreeWayQuick()
	case "AdvancedQuick":
		sorter = sorting.NewAdvancedQuick()
	default:
		panic("unsupported algs")
	}
	sorter.Sort(a, func(i, j int) bool { return a[i] < a[j] })
	consume := time.Since(start)
	utils.AssertF(sorter.IsSorted(a, func(i, j int) bool { return a[i] < a[j] }), "%s is unsorted.", alg)
	return consume
}

func (c *SortCompare) timeRandomInput(alg string, N int, T int) time.Duration {
	var total time.Duration
	a := make([]float64, N)
	for t := 0; t < T; t++ {
		for i := 0; i < N; i++ {
			a[i] = rand.Float64() + float64(rand.Intn(100))
		}
		total += c.time(alg, a)
	}
	return total
}

func main() {
	t := new(SortCompare)
	N := 1000000
	T := 10
	for _, alg := range []string{"Merge", "AdvancedMerge", "MergeBU", "Quick", "ThreeWayQuick", "AdvancedQuick"} {
		consume := t.timeRandomInput(alg, N, T)
		utils.StdOut.Println(alg, consume)
	}
}
