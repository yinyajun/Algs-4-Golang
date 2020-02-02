package main

import "fmt"

func main() {
	var a int = 842432
	mask := (1 << 8) - 1
	fmt.Println(mask)
	fmt.Printf("%032b\n", -120)
	fmt.Printf("%032b\n", a)
	fmt.Printf("%032b\n", a>>8)
	fmt.Printf("%032b\n", (a>>8)&mask)
	fmt.Printf("%032b\n", a>>(8*2))
	fmt.Printf("%032b\n", a>>(8*2)&mask)
}
