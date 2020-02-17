package string

/**
* LSD radix sort
*
*    - Sort a String[] array of n extended ASCII strings (R = 256), each of length w.
*
*    - Sort an int[] array of n 32-bit integers, treating each integer as
*      a sequence of w = 4 bytes (R = 256).
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

const BITS = 32
const BITS_PER_BYTE = 8

type LSD struct{}

// Rearranges the array of w-character strings in ascending order.
func (m *LSD) SortString(a []string, w int) {
	n := len(a)
	R := 256 // extended ASCII alphabet size
	aux := make([]string, n)

	for d := w - 1; d >= 0; d-- {
		// sort by key-indexed counting on dth character

		// compute frequency counts
		count := make([]int, R+1)
		for i := 0; i < n; i++ {
			count[a[i][d]+1]++
		}

		// compute cumulates
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// move data
		for i := 0; i < n; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}

		// copy back
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}

//// Rearranges the array of 32-bit integers in ascending order.
//func (m *LSD) SortInt(a []int) {
//	R := 1 << BITS_PER_BYTE
//	MASK := R - 1
//	w := BITS / BITS_PER_BYTE // each int is 4 bytes
//
//	n := len(a)
//	aux := []int{}
//
//	for d := 0; d < w; d++ {
//
//		// compute frequency count
//		count := make([]int, R+1)
//		for i := 0; i < n; i++ {
//			c := a[i] >> (d * BITS_PER_BYTE) & MASK
//			count[c+1]++
//		}
//
//		// compute cumulates
//		for r := 0; r < R; r++ {
//			count[r+1] += count[r]
//		}
//
//		// move data
//		for i := 0; i < n; i++ {
//			c := a[i] >> (d * BITS_PER_BYTE) & MASK
//			aux[count[c]] = a[i]
//			count[c]++
//		}
//
//		// copy back
//		for i := 0; i < n; i++ {
//			a[i] = aux[i]
//		}
//	}
//}
