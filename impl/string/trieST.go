/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/9/8 09:52
 */

package main

import (
	"Algs-4-Golang/abstract"
	"Algs-4-Golang/impl/fundamentals"
)

type TrieST struct {
	R    int
	root *Node
}

type Node struct {
	val  interface{}
	next []*Node
}

func NewNode(r int) *Node {
	return &Node{
		val:  nil,
		next: make([]*Node, r),
	}
}

func (t *TrieST) Get(key string) interface{} {
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

// 返回以x作为根节点的子字典树中与key相关的结束节点
func (t *TrieST) get(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}
	if len(key) == d {
		return x
	}
	c := int(key[d]) // 找到第d个字符串所对应的单词查找树
	return t.get(x.next[c], key, d+1)
}

func (t *TrieST) Put(key string, val interface{}) {
	t.root = t.put(t.root, key, val, 0)
}

// 返回x作为根节点的子字典树中，存入key相关的val后更新的子字典树
func (t *TrieST) put(x *Node, key string, val interface{}, d int) *Node {
	if x == nil {
		x = NewNode(t.R)
	}
	if len(key) == d {
		x.val = val
		return x
	}
	c := int(key[d])
	x.next[c] = t.put(x.next[c], key, val, d+1)
	return x
}

func (t *TrieST) Keys() abstract.Iterator { return t.KeysWithPrefix("") }

func (t *TrieST) KeysWithPrefix(pre string) abstract.Iterator {
	q := fundamentals.NewLinkedQueue()
	t.collect(t.get(t.root, pre, 0), pre, q)
	return q.Iterate()
}

func (t *TrieST) collect(x *Node, pre string, q abstract.Queue) {
	if x == nil {
		return
	}
	if x.val != nil {
		q.Enqueue(pre)
	}
	for c := 0; c < t.R; c++ {
		t.collect(x.next[c], pre+string(uint8(c)), q)
	}
}

func (t *TrieST) KeysThatMatch(pat string) abstract.Iterator {
	q := fundamentals.NewLinkedQueue()
	t.collect2(t.root, "", pat, q)
	return q.Iterate()
}

func (t *TrieST) collect2(x *Node, pre, pat string, q abstract.Queue) {
	if x == nil {
		return
	}
	if len(pre) == len(pat) {
		if x.val != nil {
			q.Enqueue(pre)
		}
		return
	}
	// len(pat) > len(pre)
	for c := 0; c < t.R; c++ {
		if pat[len(pre)] == uint8("."[0]) || pat[len(pre)] == uint8(c) {
			t.collect2(x.next[c], pre+string(uint8(c)), pat, q)
		}
	}
}

func (t *TrieST) LongestPrefixOf(s string) string {
	length := t.search(t.root, s, 0, 0)
	return string(s[:length])
}

// length记录了当前所找到的匹配s的前缀长度
func (t *TrieST) search(x *Node, s string, d, length int) int {
	if x == nil {
		return length
	}
	if x.val != nil { // 遇到key时更新length
		length = d
	}
	if d == len(s) {
		return length
	}
	c := int(s[d])
	return t.search(x.next[c], s, d+1, length)
}

// 删除操作
// 先查找，知道到键所对应的节点，并将其val=nil
// 如果该节点的所有链接均为空，需要删除当前节点；同时如果删去当前节点，使其父节点所有连接为空，也需要继续删除父节点
func (t *TrieST) delete(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		x.val = nil
	} else {
		c := int(key[d])
		x.next[c] = t.delete(x.next[c], key, d+1)
	}

	if x.val != nil {
		return x
	}

	for c := 0; c < t.R; c++ {
		if x.next[uint8(c)] != nil {
			return x
		}
	}
	return nil
}

func (t *TrieST) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

//func main() {
//	t := &TrieST{R: 256}
//	t.Put("lcc", 1)
//	t.Put("fsaky87@", 2)
//	fmt.Println(t.Get("lcc"))
//	fmt.Println(t.Get("lc"))
//	utils.PrintIterator(t.Keys())
//	utils.PrintIterator(t.KeysWithPrefix("l"))
//	println("______________")
//	utils.PrintIterator(t.KeysThatMatch("l.c"))
//	t.Put("se", 1)
//	t.Put("sea", 1)
//	fmt.Println(t.LongestPrefixOf("seab"))
//}
