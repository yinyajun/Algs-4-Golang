/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/9/10 17:01
 */

package main

import "fmt"

//当前节点t有fail指针，其fail指针所指向的节点和t所代表的字符是相同的
type node struct {
	fail  *node
	next  []*node
	count int
}

func newNode() *node {
	return &node{next: make([]*node, 256)}
}

func put(x *node, str string) {
	for i := 0; i < len(str); i++ {
		if x.next[int(str[i])] == nil {
			x.next[int(str[i])] = newNode()
		}
		x = x.next[int(str[i])]
	}
	x.count++
}

type ACAutomation struct {
	head, tail int
	q          [50000]*node
}

// Fail指针用BFS来求得
// 对于直接与根节点相连的节点来说
//		如果这些节点失配，他们的Fail指针直接指向root即可
// 其他节点其Fail指针求法如下：
// 		假设当前节点为father，其孩子节点记为child。
// 		求child的Fail指针时，首先我们要找到其father的Fail指针所指向的节点,假如是t的话，
// 		我们就要看t的孩子中有没有和child节点所表示的字母相同的节点，
// 		如果有的话，这个节点就是child的fail指针，如果发现没有，则需要找father->fail->fail这个节点，
// 		然后重复上面过程，如果一直找都找不到，则child的Fail指针就要指向root。

// 核心原理：类似于kmp中的，已经探索过的后缀是其他可能的前缀
func (a *ACAutomation) buildACAutomation(root *node) {
	a.q[a.head] = root
	a.head++

	for a.head != a.tail {
		temp := a.q[a.tail] // 设置出队节点的孩子节点的fail指针
		a.tail++

		for c := 0; c < 256; c++ {
			if temp.next[c] != nil {
				if temp == root { // 直接和root相连的节点，它们的fail指针指向root
					temp.next[c].fail = root
				} else { // 其他节点：根据父节点的fail指针来设置子节点的fail指针
					p := temp.fail
					for p != nil {
						if p.next[c] != nil {
							temp.next[c].fail = p.next[c]
							break
						}
						p = p.fail // father.fail.fail...
					}
					if p == nil { // 没找到，直接将子节点fail指针指向root
						temp.next[c].fail = root
					}
				}
				a.q[a.head] = temp.next[c]
				a.head++
			}
		}
	}
}

func (a *ACAutomation) query2(root *node, str string) int {
	var cnt int
	walkNode := root
	var nextNode *node

	for _, char := range str {
		for {
			nextNode = walkNode.next[char]
			if nextNode == nil {
				if walkNode.fail != nil {
					walkNode = walkNode.fail
					continue
				} else {
					walkNode = root
					break
				}
			} else { // match
				walkNode = nextNode
				temp := walkNode
				for temp != root && temp.count != -1 {
					cnt += temp.count
					temp.count = -1
					temp = temp.fail
				}
			}
			break
		}
	}
	return cnt
}

func (a *ACAutomation) query(root *node, str string) int {
	var cnt int
	p := root

	for i := 0; i < len(str); i++ {
		for p.next[int(str[i])] == nil && p != root { // 非根节点发生了失配str[i]
			p = p.fail
		}
		// p.next[index]!=nil || p == root
		p = p.next[int(str[i])]
		if p == nil {
			p = root
		}
		temp := p
		for temp != root && temp.count != -1 {
			cnt += temp.count
			temp.count = -1
			temp = temp.fail
		}
	}
	return cnt
}

func (a *ACAutomation) query3(root *node, str string) int {
	var cnt int
	p := root
	for i := 0; i < len(str); i++ {
		k := p.next[int(str[i])]
		if k == nil {
			k = root
		}
		for k != root && k.count != -1 {
			cnt += k.count
			k.count = -1
			k = k.fail
		}
		p = p.next[int(str[i])]
	}
	return cnt
}

func main() {
	root := newNode()
	put(root, "say")
	put(root, "she")
	put(root, "sher")
	put(root, "her")
	ac := &ACAutomation{}
	ac.buildACAutomation(root)
	str := "yasherhs"
	//fmt.Println(ac.query(root, str))
	fmt.Println(ac.query3(root, str))
}
