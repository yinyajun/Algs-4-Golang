package priorityQueue

import . "util"

type maxPQ interface {
	Insert(Key)
	Max() Key
	DelMax() Key
	IsEmpty() bool
	Size() int
}

type minPQ interface {
	Insert(Key)
	Min() Key
	DelMin() Key
	IsEmpty() bool
	Size() int
}