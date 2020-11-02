/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 20:35
 */

package utils

import (
	"bytes"
	"testing"
)

type tempStruct struct {
	a int
	b string
	c *tempStruct
}

func TestNewStdOut(t *testing.T) {
	ssim := &tempStruct{a: 2, b: "ok", c: nil}
	cases := []struct {
		a      []interface{}
		wanted string
	}{
		{[]interface{}{123}, "123"},
		{[]interface{}{243.543}, "243.543"},
		{[]interface{}{"yajun"}, "yajun"},
		{[]interface{}{243.543, 432}, "243.543 432"},
		{[]interface{}{ssim}, "&{2 ok <nil>}"},
	}
	for idx, c := range cases {
		testName := StdOut.Sprintf("case%d", idx)
		t.Run(testName, func(t *testing.T) {
			w := bytes.NewBufferString("")
			out := NewStdOut(w)
			out.Print(c.a...)
			if w.String() != c.wanted {
				t.Errorf("Got: %v, Wanted: %v", w.String(), c.wanted)
			}
		})
	}
}
