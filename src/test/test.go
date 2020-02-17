package main

import (
	"fmt"
	"util"
)

func main() {
	a := []string{"a", "c", "d", "e"}
	util.ShuffleStr(a)
	fmt.Println(a)
}
