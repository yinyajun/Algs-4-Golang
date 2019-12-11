package priorityQueue

import . "util"

type MaxPQ interface {
	insert(Key)
	max() Key
	delMax() Key
	isEmpty() bool
	size() int
}
