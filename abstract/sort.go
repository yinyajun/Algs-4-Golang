/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/3 08:28
 */

package abstract

type Sorter interface {
	Sort(slice interface{})
	Less(i, j int) bool
	Exch(i, j int)
	Show(slice interface{})
	IsSorted(slice interface{}) bool
}

//Lack of generic type, use this to index value in a interface(slice type)
type Indexer func(i int) interface{}
