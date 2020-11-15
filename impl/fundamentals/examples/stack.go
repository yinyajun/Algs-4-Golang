/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 19:15
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
	"Algs-4-Golang/utils"
)

// go run impl/fundamentals/examples/stack.go LinkedStack < data/tobe.txt
// +-------------+
// | LinkedStack |
// +-------------+
// to be not that or be ( 2 left on stack)

var s abstract.Stack

const (
	ResizeArrayStack = "ResizingArrayStack"
	LinkedStack      = "LinkedStack"
)

func init() {
	utils.Arg0 = utils.Flag.Arg(0, LinkedStack)
	utils.PrintInBox(utils.Arg0)
}

func initStack(args ...interface{}) {
	typ := args[0]
	switch typ {
	case ResizeArrayStack:
		s = fundamentals.NewResizingArrayStack()
	case LinkedStack:
		s = fundamentals.NewLinkedStack()
	}
}

func main() {
	initStack(utils.Arg0)
	for utils.StdIn.HasNext() {
		item := utils.StdIn.ReadString()
		if item != "-" {
			s.Push(item)
		} else if !s.IsEmpty() {
			utils.StdOut.Print(s.Pop(), " ")
		}
	}
	utils.StdOut.Println("(", s.Size(), "left on stack)")
}
