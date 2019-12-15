package main

import (
	"strconv"
	"fmt"
)

func main() {

	f,_:= strconv.ParseFloat("3.1415926", 64)
	fmt.Println(f)
}