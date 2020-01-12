package util

import "fmt"

type sliceIterator struct {
	c    chan interface{}
	init int
	cur  int
	list []interface{}
}

func NewListIterator(slice []interface{}) *sliceIterator {
	it := &sliceIterator{}
	it.init = 0
	it.cur = 0
	it.list = slice
	it.c = make(chan interface{})

	go func() {
		for it.cur < len(slice) {
			it.c <- slice[it.cur]
			it.cur ++
		}
		close(it.c)
	}()
	return it
}

func (s *sliceIterator) Reset() {
	s.cur = s.init
	s.c = make(chan interface{})
	go func() {
		for s.cur < len(s.list) {
			s.c <- s.list[s.cur]
			s.cur ++
		}
		close(s.c)
	}()
}

func main() {
	list := []interface{}{}
	list = append(list, 5)
	list = append(list, 6)
	list = append(list, 7)

	it := NewListIterator(list)
	for e := range it.c {
		fmt.Println(e)
	}

	it.Reset()
	for e := range it.c {
		fmt.Println(e)
	}

}
