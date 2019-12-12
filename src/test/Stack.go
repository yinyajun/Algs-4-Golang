package main

import (
	"os"
	"fmt"

	. "algs4/stack"
	. "util"
)

func main() {
	s := NewStack()
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
	PrintIterator(s)
}
