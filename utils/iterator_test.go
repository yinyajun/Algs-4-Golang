/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/01 16:56
 */

package utils

import (
	"testing"
)

func TestNewRangeIterator(t *testing.T) {
	cases := []struct {
		first, last, interval int
		reverse               bool
		want                  []interface{}
	}{
		{2, 5, 1, false, []interface{}{2, 3, 4}},
		{2, 1, 1, false, []interface{}{}},
		{2, 10, 2, false, []interface{}{2, 4, 6, 8}},
		{2, 10, 2, true, []interface{}{10, 8, 6, 4}},
	}
	for idx, c := range cases {
		testName := StdOut.Sprintf("case%d", idx)
		t.Run(testName, func(t *testing.T) {
			it := NewRangeIterator(c.first, c.last, c.interval, c.reverse)
			idx := 0
			for it.First(); it.HasNext(); {

				number := it.Next()
				if c.want[idx] != number {
					t.Errorf("got: %v, want: %v", number, c.want[idx])
				}
				idx++
			}
			if idx != len(c.want) {
				t.Errorf("got length: %d, want length: %d", idx, len(c.want))
			}
		})
	}
}

func TestNewCircularIterator(t *testing.T) {
	cases := []struct {
		first, last, interval int
		reverse               bool
		mod, repeat           int
		want                  []interface{}
	}{
		{2, 2, 1, false, 6, 1, []interface{}{2, 3, 4, 5, 0, 1}}, // last:2 occur once
		{2, 2, 1, true, 6, 1, []interface{}{2, 1, 0, 5, 4, 3}},  // last:2 occur once
		{2, 2, 1, false, 6, 0, []interface{}{}},                 // last:2 never occur
		{2, 2, 1, false, 0, 0, []interface{}{}},

		{2, 5, 1, false, 0, 0, []interface{}{2, 3, 4}},
		{2, 5, 1, false, 5, 0, []interface{}{2, 3, 4}},                         // last:5 never occur
		{2, 5, 1, false, 6, 0, []interface{}{2, 3, 4}},                         // last:5 never occur
		{2, 5, 1, true, 6, 0, []interface{}{5, 4, 3}},                          // last:2 never occur(reverse)
		{2, 5, 1, false, 6, 1, []interface{}{2, 3, 4, 5, 0, 1, 2, 3, 4}},       // last:5 occur once
		{2, 1, 1, false, 6, 0, []interface{}{2, 3, 4, 5, 0}},                   // last:1 never occur
		{2, 1, 1, false, 6, 1, []interface{}{2, 3, 4, 5, 0, 1, 2, 3, 4, 5, 0}}, // last:1 occur once
	}
	for idx, c := range cases {
		testName := StdOut.Sprintf("case%d", idx)
		t.Run(testName, func(t *testing.T) {
			it := NewCircularIterator(c.first, c.last, c.interval, c.mod, c.repeat, c.reverse)
			idx := 0
			for it.First(); it.HasNext(); {
				number := it.Next()
				if c.want[idx] != number {
					t.Errorf("got: %v, want: %v", number, c.want[idx])
				}
				idx++
			}
			if idx != len(c.want) {
				t.Errorf("got length: %d, want length: %d", idx, len(c.want))
			}
		})
	}
}

func TestNewArrayIterator(t *testing.T) {
	slice1 := []interface{}{0, 1, 2, 3, 4}
	cases := []struct {
		slice             []interface{}
		first, size       int
		circular, reverse bool
		want              []interface{}
	}{
		{slice1, 2, 2, false, false, []interface{}{2, 3}},
		{slice1, 2, 3, false, false, []interface{}{2, 3, 4}},
		{slice1, 2, 7, true, false, []interface{}{2, 3, 4, 0, 1, 2, 3}},
	}
	for idx, c := range cases {
		testName := StdOut.Sprintf("case%d", idx)
		t.Run(testName, func(t *testing.T) {
			it := NewArrayIterator(c.slice, c.first, c.size, c.circular)
			idx := 0
			for it.First(); it.HasNext(); {
				number := it.Next()
				if c.want[idx] != number {
					t.Errorf("got: %v, want: %v", number, c.want[idx])
				}
				idx++
			}
			if idx != len(c.want) {
				t.Errorf("got length: %d, want length: %d", idx, len(c.want))
			}
		})
	}
}
