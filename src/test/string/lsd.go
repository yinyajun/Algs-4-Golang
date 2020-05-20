package main

import (
	string2 "algs4/string"
	"fmt"
	"os"
	"util"
)

/**
*
* $ go run src/test/lsd.go < data/words3.txt
* all
* bad
* bed
* ...
* wee
* yes
* yet
* zoo
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := util.NewIn(os.Stdin)
	a := []string{}
	for in.HasNext() {
		a = append(a, in.ReadString())
	}
	n := len(a)

	// check that strings have fixed length
	w := len(a[0])
	for i := 0; i < n; i++ {
		if len(a[i]) != w {
			panic("Strings must have fixed length")
		}
	}

	// sort the strings
	lsd := &string2.LSD{}
	lsd.SortString(a, w)

	// print results
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}
