package main

import "fmt"

type a struct {
	b int
}

func (m *a) String() string {
	return "1"
}
func main() {
	c := &a{}
	fmt.Println(c)
}
