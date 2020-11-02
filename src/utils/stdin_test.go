/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:57
 */

package utils

import (
	"strings"
	"testing"
)

func TestStdIn_ReadString(t *testing.T) {
	r := strings.NewReader("to be or not to be")
	in := NewStdIn(r, "words")
	ans := []string{"to", "be", "or", "not", "to", "be"}
	idx := 0
	for in.HasNext() {
		k := in.ReadString()
		if ans[idx] != k {
			t.Errorf("read: [%v], acutually: [%v]", k, ans[idx])
		}
		idx++
	}
	if idx != len(ans) {
		t.Errorf("read length larger than actual string slice")
	}
}

func TestStdIn_ReadInt(t *testing.T) {
	r := strings.NewReader("1 2 3 4 5 6")
	in := NewStdIn(r, "words")
	ans := []int{1, 2, 3, 4, 5, 6}
	idx := 0
	for in.HasNext() {
		k := in.ReadInt()
		if ans[idx] != k {
			t.Errorf("read: [%v], acutually: [%v]", k, ans[idx])
		}
		idx++
	}
	if idx != len(ans) {
		t.Errorf("read length larger than actual string slice")
	}
}

func TestStdIn_ReadLine(t *testing.T) {
	r := strings.NewReader("to be or\n not to be\n")
	in := NewStdIn(r, "line")
	ans := []string{"to be or", " not to be"}
	idx := 0
	for in.HasNext() {
		k := in.ReadLine()
		if ans[idx] != k {
			t.Errorf("read: [%v], acutually: [%v]", k, ans[idx])
		}
		idx++
	}
	if idx != len(ans) {
		t.Errorf("read length larger than actual string slice")
	}
}
