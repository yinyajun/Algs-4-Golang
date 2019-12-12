package util

import (
	"bufio"
	"io"
	"strconv"
)

type In struct {
	*bufio.Scanner
	buf  []string
	rpos int
}

func NewIn(r io.Reader) *In {
	return NewInWithSplitFunc(r, bufio.ScanWords)
}

func NewInWithSplitFunc(r io.Reader, split bufio.SplitFunc) *In {
	s := bufio.NewScanner(r)
	s.Split(split)
	m := &In{s, []string{}, 0}
	return m
}

func (m *In) HasNext() bool {
	if m.Scan() {
		m.buf = append(m.buf, m.Text())
		return true
	}
	return false
}

func (m *In) ReadString() string {
	defer func() { m.rpos++ }()
	if m.rpos < len(m.buf) {
		return m.buf[m.rpos]
	}
	if m.HasNext() {
		return m.buf[m.rpos]
	}
	panic("no more input to read")
}

func (m *In) ReadInt() int {
	s := m.ReadString()
	i, _ := strconv.Atoi(s)
	return i
}