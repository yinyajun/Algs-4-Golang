package main

import (
	"fmt"
	"os"

	. "algs4/bag"
	. "util"
)

/**
* $ go run src/test/bag.go < data/tobe.txt
* size of bag =  14
* is
* -
* -
* -
* that
* -
* -
* be
* -
* to
* not
* or
* be
* to
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	bag := NewBag()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		bag.Add(in.ReadString())
	}
	fmt.Println("size of bag = ", bag.Size())
	PrintIterators(bag.Iterate())
}
