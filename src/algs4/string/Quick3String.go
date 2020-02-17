package string

import (
	"reflect"
	"util"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type Quick3String struct {
	CUTOFF int
}

func NewQuick3String() *Quick3String {
	return &Quick3String{15}
}

func (q *Quick3String) Sort(a []string) {
	util.ShuffleStr(a)
	q.sort(a, 0, len(a)-1, 0)
	if !q.isSorted(a) {
		panic("Sort error.")
	}
}

// return dth character of s, -1 if d = length of string
func (q *Quick3String) charAt(s string, d int) int {
	if d < 0 || d > len(s) {
		panic("invalid d")
	}
	if d == len(s) {
		return -1
	}
	return int(s[d])
}

// 3-way string quicksort a[lo..hi] starting at dth character
func (q *Quick3String) sort(a []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+q.CUTOFF {
		q.insertion(a, lo, hi, d)
		return
	}

	lt := lo // [lo...lt) < v
	gt := hi // (gt, hi] > v
	v := q.charAt(a[lo], d)
	i := lo + 1
	swap := reflect.Swapper(a)
	for i <= gt {
		t := q.charAt(a[i], d)
		if t < v {
			swap(i, lt)
			i++
			lt++
		} else if t > v {
			swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	q.sort(a, lo, lt-1, d)
	if v >= 0 {
		q.sort(a, lt, gt, d+1)
	}
	q.sort(a, gt+1, hi, d)
}

// sort from a[lo] to a[hi], starting at the dth character
func (q *Quick3String) insertion(a []string, lo, hi, d int) {
	for i := lo + 1; i <= hi; i++ {
		for j := i; j > lo && q.less(a[j], a[j-1], d); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

// is v less than w, starting at character d
func (q *Quick3String) less(v, w string, d int) bool {
	// assume v.substring(0, d).equals(w.substring(0, d))
	for i := d; i < len(v) && i < len(w); i++ {
		if v[i] < w[i] {
			return true
		}
		if v[i] > w[i] {
			return false
		}
	}
	return len(v) < len(w)
}

func (q *Quick3String) isSorted(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
