package util

import (
	"math/rand"
	"reflect"
	"time"
)

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func ShuffleStr(a []string) {
	// knuth-Shuffle
	r := rand.New(rand.NewSource(time.Now().Unix()))
	swap := reflect.Swapper(a) // swapper function
	for n := len(a); n > 0; n-- {
		swap(r.Intn(n), n-1)
	}
}
