/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:28
 */

package abstract

type Sorter interface {
	Sort(slice interface{}, less func(i, j int) bool)
	Exch(i, j int)
	IsSorted(slice interface{}, less func(i, j int) bool) bool
}
