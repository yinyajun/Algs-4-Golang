/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/9/8 15:19
 */

package main

type KMP struct {
	DFA [][]int // row: 字符集；column：状态数（m+1）
	// 总共有【0，1，...,m】共m+1个状态，其中m状态为停机状态
}

// 每个状态代表pattern能匹配到第几个字母，当匹配到停机状态就说明匹配成功
// 所有非停机状态[0...m-1]，都会有状态转移，所以有m个状态需要状态转移

// dfa[c][j]: 当前状态j，意味着已经匹配了j个字符pat[0...j)，下一个待匹配字符为c。根据pattern[j]和c的值，会跳转到不同的状态
// * 如果匹配，会跳转到下一个状态j+1: dfa[c][j] = j+1
// * 如果失配，pattern会回退，而txt不回退，为了尽可能减少pattern的回退, 将pattern回退到重启状态x: dfa[c][j] = dfa[c][x]
//        此时pattern的前缀[0...x]和（已匹配的txt后缀+c）对齐（所以需要记录公共后缀，dfa值的另一种理解）

func NewKMP1(pattern string) *KMP {
	r := 256
	m := len(pattern)
	kmp := &KMP{DFA: make([][]int, r)}

	// init DFA
	for i := 0; i < r; i++ {
		kmp.DFA[i] = make([]int, m)
	}

	// base
	kmp.DFA[int(pattern[0])][0] = 1
	// 在状态0（未匹配任何字符）
	// * 如果遇到字符为为pattern[0]，那么能够匹配，进入状态1
	// * 遇到其他字符必然不匹配，所以仍然转移到状态0

	// construct dfa
	x := 0 // 重启状态初始化为0
	for j := 1; j < m; j++ {
		for c := 0; c < r; c++ {
			kmp.DFA[c][j] = kmp.DFA[c][x] // 默认失配，回到重启状态
		}
		kmp.DFA[int(pattern[j])][j] = j + 1 // 更新匹配状态
		x = kmp.DFA[int(pattern[j])][x]     // 更新重启状态（重启状态意味着未匹配前的最大公共前后缀，此时遇到pattern[j]，它会怎么更新？直接调用DFA[][x]即可）
	}
	return kmp
}

func (k *KMP) search(text string) int {
	n := len(text)
	m := len(k.DFA[0])

	j := 0 // pattern的初始状态
	for i := 0; i < n; i++ {
		j = k.DFA[int(text[i])][j] // 计算pattern的下一个状态
		if j == m {                // 到达停机状态
			return i - m + 1
		}
	}
	return n
}

// 构建DFA的过程和search的过程十分相似
// search在text中匹配pattern
// construct在pattern[1...]中匹配pattern（即找公共前后缀）
