/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:55
 */

package utils

import (
	"Algs-4-Golang/abstract"
)

func PrintIterator(it abstract.Iterator) {
	for it.First(); it.HasNext(); {
		StdOut.Println(it.Next())
	}
}

// 使用模板模式，作为具体的iterator的模板
// 将Ranger接口的方法延迟到子类中实现（父类中提供默认实现）
type RangeBase struct {
	first    int
	last     int
	interval int
	reverse  bool
	cur      int
	impl     abstract.Ranger // subclass
}

func (b *RangeBase) Init(first, last, interval int, reverse bool, impl abstract.Ranger) {
	b.first = first
	b.last = last
	b.interval = interval
	b.reverse = reverse
	b.impl = impl
}

func (b *RangeBase) InitFirstIndex() int { return b.first }

func (b *RangeBase) InitLastIndex() int { return b.last }

func (b *RangeBase) GetIndex(index int) int { return index }

func (b *RangeBase) GetNextIndex() int {
	if b.reverse {
		return b.cur - b.interval
	}
	return b.cur + b.interval
}

func (b *RangeBase) First() {
	b.first = b.impl.InitFirstIndex()
	b.last = b.impl.InitLastIndex()
	if b.reverse {
		b.cur = b.last
		return
	}
	b.cur = b.first
}

func (b *RangeBase) HasNext() bool {
	if b.reverse {
		return b.cur > b.first
	}
	return b.cur < b.last
}

func (b *RangeBase) Next() interface{} {
	if !b.HasNext() {
		return nil
	}
	res := b.impl.GetIndex(b.cur)
	b.cur = b.GetNextIndex()
	return res
}

// [first, last) supports circular accessing
type CircularIterator struct {
	*RangeBase
	mod    int // mod==0 implies never circular
	repeat int // occurrence of encounter (last-th index) when mod > 0
}

func NewCircularIterator(first, last, interval, mod, repeat int, reverse bool) *CircularIterator {
	it := new(CircularIterator)
	base := new(RangeBase)
	base.Init(first, last, interval, reverse, it)
	it.Init(base, mod, repeat)
	return it
}

func (c *CircularIterator) Init(base *RangeBase, mod, repeat int) {
	c.RangeBase = base
	c.mod = mod
	c.repeat = repeat
	Assert(c.mod >= 0, "invalid parameter mod")
	Assert(c.repeat >= 0, "invalid parameter repeat")
	Assert(c.mod == 0 || c.mod > 0 && c.last <= c.mod, "invalid parameter last")
}

func (c *CircularIterator) InitLastIndex() int {
	// 找到第一个大于first的last索引（同余关系下）
	// new_last = last < first ? last+mod, last
	// new_last >= last && new_last(mod mod) === last
	last := c.last
	if last < c.first {
		last += c.mod
	}
	return last + c.mod*c.repeat
}

func (c *CircularIterator) GetIndex(index int) int {
	if c.mod > 0 {
		return index % c.mod
	}
	return index
}

// [first, last)
type RangeIterator struct {
	*RangeBase
}

func NewRangeIterator(first, last, interval int, reverse bool) *RangeIterator {
	it := new(RangeIterator)
	base := new(RangeBase)
	base.Init(first, last, interval, reverse, it)
	it.RangeBase = base
	return it
}

type ArrayIterator struct {
	*RangeBase
	mod   int
	slice []interface{}
}

func NewArrayIterator(slice []interface{}, first, size int, circular bool) *ArrayIterator {
	it := new(ArrayIterator)
	base := new(RangeBase)
	base.Init(first, first+size, 1, false, it)
	it.Init(base, slice, circular)
	return it
}

func (a *ArrayIterator) Init(base *RangeBase, slice []interface{}, circular bool) {
	a.RangeBase = base
	a.slice = slice
	if circular {
		a.mod = len(slice)
	} else {
		// [first, len(slice)) include all elements if not circular
		Assert(base.last <= len(slice), "Invalid index last")
	}
}

func (a *ArrayIterator) GetIndex(index int) int {
	if a.mod > 0 {
		return index % a.mod
	}
	return index
}

func (a *ArrayIterator) Next() interface{} {
	idx := a.RangeBase.Next().(int)
	return a.slice[idx]
}

type LinkedListIterator struct {
	head abstract.Node
	cur  abstract.Node
}

func NewLinkedListIterator(node abstract.Node) *LinkedListIterator {
	it := &LinkedListIterator{}
	it.head = node
	return it
}

func (i *LinkedListIterator) First() {
	i.cur = i.head
}

func (i *LinkedListIterator) HasNext() bool {
	return i.cur != nil
}

func (i *LinkedListIterator) Next() interface{} {
	if i.HasNext() {
		res := i.cur.Key()
		i.cur = i.cur.Next()
		return res
	}
	return nil
}
