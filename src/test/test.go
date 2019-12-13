package main

import "fmt"

func main() {
	c:= make(map[string]int)
	fmt.Println(len(c))
	c["a"]=5
	fmt.Println(len(c))
}
