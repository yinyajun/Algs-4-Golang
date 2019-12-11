package main

import (
	"fmt"
	"os"

	. "algs4/graph"
	. "util"
)

func main() {
	in := NewIn(os.Stdin)
	g := NewGraphWithIn(in)
	fmt.Println(g)
}
