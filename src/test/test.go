package main

import (
	"github.com/deckarep/golang-set"
	"fmt"
)

type YourType struct {
	Name string
}

func main() {
	set := mapset.NewThreadUnsafeSetFromSlice([]interface{}{
		&YourType{Name: "Alise"},
		&YourType{Name: "Bob"},
		&YourType{Name: "John"},
		&YourType{Name: "Nick"},
	})

	var found *YourType
	it := set.Iterator()

	for elem := range it.C {
		if elem.(*YourType).Name == "John" {
			found = elem.(*YourType)
			it.Stop()
		}
	}

	fmt.Printf("Found %+v\n", found)

	// Output: Found &{Name:John}
}
