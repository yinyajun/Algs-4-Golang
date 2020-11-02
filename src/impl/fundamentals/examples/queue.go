/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 19:26
 */

package main

import (
	"abstract"
	"impl/fundamentals"
	"utils"
)

// ./run.sh src/impl/fundamentals/examples/queue.go ResizingArrayQueue < data/tobe.txt
// +--------------------+
// | ResizingArrayQueue |
// +--------------------+
// to be or not to be ( 2 left on queue)

const (
	ResizingArrayQueue = "ResizingArrayQueue"
	LinkedQueue        = "LinkedQueue"
)

var q abstract.Queue

func init() {
	utils.Arg0 = utils.Flag.Arg(0, LinkedQueue)
	utils.PrintInBox(utils.Arg0)
}

func initQueue(args ...interface{}) {
	typ := args[0]
	switch typ {
	case ResizingArrayQueue:
		q = fundamentals.NewResizingArrayQueue()
	case LinkedQueue:
		q = fundamentals.NewLinkedQueue()
	}
}

func main() {
	initQueue(utils.Arg0)
	for utils.StdIn.HasNext() {
		item := utils.StdIn.ReadString()
		if item != "-" {
			q.Enqueue(item)
		} else if !q.IsEmpty() {
			utils.StdOut.Print(q.Dequeue(), " ")
		}
	}
	utils.StdOut.Println("(", q.Size(), "left on queue)")
}
