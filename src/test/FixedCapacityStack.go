package main

import (
	"fmt"
	"os"

	. "algs4/stack"
	. "util"
)

/**
* $ go run  src/test/FixedCapacityStack.go < data/tobe.txt
* to be not that or be ( 2 left on stack)
* is
* to
*
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	s := NewFixedCapacityStrings(100)
	in := NewIn(os.Stdin)
	for in.HasNext() {
		item := in.ReadString()
		if item != "-" {
			s.Push(item)
		} else if !s.IsEmpty() {
			fmt.Print(s.Pop(), " ")
		}
	}
	fmt.Println("(", s.Size(), "left on stack)")
	PrintIterator(s.Iterate())
}
