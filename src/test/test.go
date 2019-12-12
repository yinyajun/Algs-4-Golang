package main

import (
	"os"
	"util"
	"bufio"
)

func main() {
	f, _ := os.Open("graph.go")
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	in := util.NewInWithSplitFunc(f, bufio.ScanLines)




}
