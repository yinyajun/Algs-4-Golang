package main

import (
	"algs4/fundamental"
	"flag"
	"fmt"
)

var (
	name string
	help bool
)

func main() {
	flag.StringVar(&name, "n", "", "algorithm name")
	flag.BoolVar(&help, "h", false, "need help")
	flag.Parse()
	if help || name == "" {
		flag.Usage()
		return
	}
	algorithmRoute(name)
}

func algorithmRoute(name string) {
	switch name {
	case "FixedCapacityStackOfStrings":
		fundamental.EgFCSS()
	case "FixedCapacityStack":
		fundamental.EgFCS()
	case "Stack":
		fundamental.EgStack()
	case "Queue":
		fundamental.EgQueue()
	default:
		fmt.Println(name, "not exists, check algorithm name.")
	}
}
