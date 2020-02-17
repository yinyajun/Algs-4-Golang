package main

import (
	string2 "algs4/string"
	"fmt"
	"os"
	"util"
)

/**
*
* $ go run src/test/msd.go < data/shells.txt
* are
* by
* sea
* seashells
* seashells
* sells
* sells
* she
* she
* shells
* shore
* surely
* the
* the
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := util.NewIn(os.Stdin)
	a := []string{}
	for in.HasNext() {
		a = append(a, in.ReadString())
	}
	n := len(a)

	msd := string2.NewMSD()
	msd.SortingString(a)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}
