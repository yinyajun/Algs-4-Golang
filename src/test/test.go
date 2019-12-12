package main

import "fmt"

func main() {
	const INT_MAX = int(^uint(0) >> 1)
	fmt.Println(INT_MAX)
	fmt.Println(^INT_MAX)
}