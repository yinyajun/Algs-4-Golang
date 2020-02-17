package string

/**
*  provides static methods for sorting an
*  array of extended ASCII strings or integers using MSD radix sort.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type MSD struct {
	R      int // extended ASCII alphabet size
	CUTOFF int // cutoff to insertion sort
}

func NewMSD() *MSD {
	return &MSD{256, 15}
}

// Rearranges the array of extended ASCII strings in ascending order.
func (m *MSD) SortingString(a []string) {
	n := len(a)
	aux := make([]string, n)
	m.sort(a, 0, n-1, 0, aux)
}

// return dth character of s, -1 if d = length of string
func (m *MSD) charAt(s string, d int) int {
	if d < 0 || d > len(s) {
		panic("invalid d")
	}
	if d == len(s) {
		return -1
	}
	return int(s[d])
}

// sort from a[lo] to a[hi], starting at the dth character
func (m *MSD) sort(a []string, lo, hi, d int, aux []string) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+m.CUTOFF {
		m.insertion(a, lo, hi, d)
		return
	}

	// compute frequency counts
	count := make([]int, m.R+2)
	for i := lo; i <= hi; i++ {
		c := m.charAt(a[i], d)
		count[c+2]++
	}

	// transform counts to indices
	for r := 0; r < m.R+1; r++ {
		count[r+1] += count[r]
	}

	// distribute
	for i := lo; i <= hi; i++ {
		c := m.charAt(a[i], d)
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < m.R; r++ {
		m.sort(a, lo+count[r], hi+count[r+1]-1, d+1, aux)
	}
}

// insertion sort a[lo..hi], starting at dth character
func (m *MSD) insertion(a []string, lo, hi, d int) {
	for i := lo + 1; i <= hi; i++ {
		for j := i; j > lo && m.less(a[j], a[j-1], d); j-- {
			m.exch(a, j, j-1)
		}
	}
}

// exchange a[i] and a[j]
func (m *MSD) exch(a []string, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// is v less than w, starting at character d
func (m *MSD) less(v, w string, d int) bool {
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

// todo: msd for []int
