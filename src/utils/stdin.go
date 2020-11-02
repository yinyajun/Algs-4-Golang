/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 10:57
 */

package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const (
	initState = iota
	accessState
	endState
)

type stdIn struct {
	state   int
	scanner *bufio.Scanner
}

var (
	StdIn     *stdIn
	StdInLine *stdIn
)

func init() {
	StdIn = NewStdIn(os.Stdin, "words")
	StdInLine = NewStdIn(os.Stdin, "line")
}

// tempStruct factory
func NewStdIn(r io.Reader, typ string) *stdIn {
	scanner := bufio.NewScanner(r)
	switch typ {
	case "words":
		scanner.Split(bufio.ScanWords)
	case "line":
		scanner.Split(bufio.ScanLines)
	case "byte":
		scanner.Split(bufio.ScanBytes)
	default:
		Panic("unsupported stdin mode")
	}
	return &stdIn{initState, scanner}
}

func (in *stdIn) HasNext() bool {
	for {
		switch in.state {
		case initState:
			if in.scanner.Scan() {
				in.state = accessState
			} else {
				in.state = endState
			}
		case accessState:
			return true
		case endState:
			return false
		}
	}
}

func (in *stdIn) ReadString() string {
	if in.HasNext() {
		in.state = initState
		return in.scanner.Text()
	}
	panic("stdIn error: eof")
}

func (in *stdIn) ReadInt() int {
	s := in.ReadString()
	i, _ := strconv.Atoi(s)
	return i
}

func (in *stdIn) ReadFloat() float64 {
	s := in.ReadString()
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func (in *stdIn) ReadLine() string {
	// in.scanner.split should be bufio.ScanLines
	return in.ReadString()
}
