package main

import (
	"algs4/graph"
	"fmt"
	"strings"
)

func main() {
	e := graph.NewDirectedEdge(12, 34, 5.67)
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%v\n", e))
	s.WriteString(fmt.Sprintf("%T\n", e))
	s.WriteString(fmt.Sprintf("%s\n", e))
	fmt.Println(s.String())
}
