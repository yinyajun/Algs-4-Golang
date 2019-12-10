package util

import (
	"strings"
	"fmt"
	"testing"
)

func TestNewIn(t *testing.T) {
	s := strings.NewReader("to be\nor not to be")
	in := NewIn(s)
	ans := []string{"to", "be", "or", "not", "to", "be"}
	idx := 0
	for in.HasNext() {
		k := in.ReadString()
		if ans[idx] != k {
			fmt.Println("read: [%v], acutually: [%v]", k, ans[idx])
		}
		idx ++
	}
	if idx != len(ans) {
		fmt.Println("read length larger than actual string array")
	}
}
