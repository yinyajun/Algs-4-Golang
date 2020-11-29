/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/27 10:41
 */

package utils

// 类似于索引优先队列中的pq和qp数组
// 因为golang中没有泛型，排序任务需要用reflect来实现，而reflect性能较差
// 通过该函数，构造原数组的正向索引forward和反向索引reverse的数组，直接对正向索引forward数组操作而不是对原数组操作
// 通过这样的操作，极大加快排序的速度
func BuildIndexSlice(length int) []int {
	indexes := make([]int, length)
	for i := 0; i < length; i++ {
		indexes[i] = i // idx -> slice_idx
	}
	return indexes
}

// 已经有排好序的正向索引数组fwd，那么根据fwd数组，通过swap函数，更新原数组
// 具体而言，在原数组第i个位置上，放上fwd[i]对应的原数组值
// 在排定过程中，原数组会改变，需要同时更新fwd和reverse数组来保证正确性
func SortByIndex(indexes []int, swap func(i, j int)) {
	length := len(indexes)
	reverse := make([]int, length)
	// init reverse indexes
	for i := 0; i < length; i++ {
		reverse[indexes[i]] = i // slice_idx -> idx
	}

	for i := 0; i < length; i++ {
		iIdx := indexes[i]
		swap(i, iIdx)
		indexes[reverse[i]], indexes[reverse[iIdx]] = indexes[reverse[iIdx]], indexes[reverse[i]]
		reverse[iIdx], reverse[i] = reverse[i], reverse[iIdx]
	}
}
