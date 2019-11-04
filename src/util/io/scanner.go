package io

import (
	"bufio"
	"fmt"
	"os"
)

type In struct {
	*bufio.Scanner
}

func SplitFunc(name string) bufio.SplitFunc {
	switch name {
	case "words":
		return bufio.ScanWords
	default:
		return bufio.ScanWords
	}
}

func NewIn(split bufio.SplitFunc) *In {
	s := bufio.NewScanner(os.Stdin)
	s.Split(split)
	return &In{s}
}

func (m *In) IsEmpty() bool {
	return !m.Scan()
}

func (m *In) ReadString() string {
	return m.Text()
}

func (m *In) ReadAllStrings() []string {
	ret := []string{}
	for !m.IsEmpty() {
		ret = append(ret, m.ReadString())
	}
	return ret
}

func main() {
	s := NewIn(bufio.ScanWords)
	ret := []string{}
	for !s.IsEmpty() {
		ret = append(ret, s.ReadString())
	}
	fmt.Println(ret)
}
