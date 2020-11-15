/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 16:55
 */

package utils

import (
	"testing"
)

func TestGreat(t *testing.T) {
	cases := []struct {
		a, b   interface{}
		wanted bool
	}{
		{1, 2, false},
		{"1", "2", false},
		{1.1, 1.2, false},
	}
	for idx, c := range cases {
		testName := StdOut.Sprintf("case%d", idx)
		t.Run(testName, func(t *testing.T) {
			if Great(c.a, c.b) != c.wanted {
				t.Errorf("got: %v, wanted: %v", !c.wanted, c.wanted)
			}
		})
	}
}
