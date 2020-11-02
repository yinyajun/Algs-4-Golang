/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 15:43
 */

package main

import (
	"abstract"
	"impl/fundamentals"
	"math"
	"strconv"
	"utils"
)

// Dijkstra双栈算数表达式求值
// ---------------------------------------------------------
// ./run.sh src/impl/fundamentals/examples/evaluate.go
// ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
// 101

func main() {
	var ops abstract.Stack
	var values abstract.Stack
	ops = fundamentals.NewLinkedStack()
	values = fundamentals.NewLinkedStack()
	for utils.StdIn.HasNext() {
		s := utils.StdIn.ReadString()
		switch s {
		case "(":
		case "+":
			ops.Push("+")
		case "-":
			ops.Push("-")
		case "*":
			ops.Push("*")
		case "/":
			ops.Push("/")
		case "sqrt":
			ops.Push("sqrt")
		case ")":
			// 弹出运算符和操作数，计算结果并压入操作数栈
			v := values.Pop().(float64)
			op := ops.Pop().(string)
			switch op {
			case "+":
				v = values.Pop().(float64) + v
			case "-":
				v = values.Pop().(float64) - v
			case "*":
				v = values.Pop().(float64) * v
			case "/":
				v = values.Pop().(float64) * v
			case "sqrt":
				v = math.Sqrt(v)
			}
			values.Push(v)
		default:
			// value
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				utils.Panic("invalid value")
			}
			values.Push(v)
		}
	}
	utils.StdOut.Println(values.Pop())
}
