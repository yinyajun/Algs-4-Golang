package main

import (
	"os"
	"fmt"

	. "algs4/bag"
	. "util"
)

func main() {
	bag := NewBag()
	in := NewIn(os.Stdin)
	for in.HasNext() {
		bag.Add(in.ReadString())
	}
	fmt.Println("size of bag = ", bag.Size())
	for _, i := range bag.Iterator() {
		fmt.Println(i)
	}
}