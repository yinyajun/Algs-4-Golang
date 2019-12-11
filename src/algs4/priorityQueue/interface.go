package priorityQueue

import . "util"

type MaxPQ interface {
	Insert(Key)
	Max() Key
	DelMax() Key
	IsEmpty() bool
	Size() int
}
