package stack

/**
*
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
*/

type FixedCapacityStack struct {
	a []interface{}
	N int
}

func NewFixedCapacityStrings(cap int) *FixedCapacityStack {
	return &FixedCapacityStack{
		a: make([]interface{}, cap),
	}
}

func (m *FixedCapacityStack) IsEmpty() bool {
	return m.N == 0
}

func (m *FixedCapacityStack) Size() int {
	return m.N
}

func (m *FixedCapacityStack) Push(item interface{}) {
	m.a[m.N] = item
	m.N++
}
func (m *FixedCapacityStack) Pop() interface{} {
	m.N--
	ret := m.a[m.N]
	return ret
}


