/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:58
 */

package utils

import (
	"math/rand"
	"reflect"
)

var Random *random

func init() {
	Random = NewRandom(123)
}

type random struct {
	*rand.Rand
	seed int64
}

func NewRandom(seed int64) *random {
	r := &random{seed: seed}
	r.Rand = rand.New(rand.NewSource(r.seed))
	return r
}

func (r *random) SetSeed(seed int64) { r.seed = seed }

func (r *random) Shuffle(a interface{}) {
	rv := reflect.ValueOf(a)
	swap := reflect.Swapper(a)
	for n := rv.Len(); n > 0; n-- {
		swap(r.Intn(n), n-1)
	}
}
