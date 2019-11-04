package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	*bufio.Scanner
	eof bool
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
	return &In{s, false}
}

func (m *In) IsEmpty() bool {
	return m.eof
}

func (m *In) ReadString() string {
	if m.Scan() {
		m.eof = false
	} else {
		m.eof = true
	}
	return m.Text()
}

func (m *In) ReadInt() int {
	s := m.ReadString()
	i, _ := strconv.Atoi(s)
	return i
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
