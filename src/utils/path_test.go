/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 13:50
 */

package utils

import (
	"testing"
)

func TestTruncateSplit(t *testing.T) {
	cases := []struct {
		fileName string
		level    int
		want     string
	}{
		{"/a/b/c", 2, "b/c"},
		{"a/b/c", 2, "b/c"},
		{"/a/b/c", 5, "a/b/c"},
		{"s/a/b/c", 5, "s/a/b/c"},
	}
	for _, c := range cases {
		testName := StdOut.Sprintf("case[%s, %d]", c.fileName, c.level)
		t.Run(testName, func(t *testing.T) {
			_, fileName := TruncateSplit(c.fileName, c.level)
			if fileName != c.want {
				t.Errorf("got: %s, want: %s", fileName, c.want)
			}
		})
	}
}
