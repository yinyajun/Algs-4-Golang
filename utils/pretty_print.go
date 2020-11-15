/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 17:30
 */

package utils

import (
	"strings"
)

const (
	CornerStr = "+"
)

// 将str用空格填充到length长度
func padString(str string, length int) string {
	sb := new(strings.Builder)
	sb.WriteString(str)
	sb.WriteString(fill(" ", length-len(str)))
	return sb.String()
}

// 返回num个str组成的string
func fill(str string, num int) string {
	sb := new(strings.Builder)
	for i := 0; i < num; i++ {
		sb.WriteString(str)
	}
	return sb.String()
}

// 找到二维string数组中的每个string的最大值，作为表格的最大宽度
func findMaxLength(infosArray ...[]string) int {
	length := 0
	for _, infos := range infosArray {
		for i := range infos {
			if len(infos[i]) > length {
				length = len(infos[i])
			}
		}
	}
	return length
}

func printBody(infos []string, length int) {
	for _, info := range infos {
		StdOut.Printf("| %s |\n", padString(info, length))
	}
}

func PrintInBox(infos ...string) {
	length := findMaxLength(infos)
	line := CornerStr + fill("-", length+2) + CornerStr
	StdOut.Println(line)
	printBody(infos, length)
	StdOut.Println(line)
}

// 以表格形式输出，输入是二维string数组，每个一维的string数组在一个格子内
func PrintInTable(infosArray ...[]string) {
	length := findMaxLength(infosArray...)
	line := CornerStr + fill("-", length+2) + CornerStr
	StdOut.Println(line)
	for _, infos := range infosArray {
		printBody(infos, length)
		StdOut.Println(line)
	}
}
