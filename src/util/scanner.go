package util

import (
	"bufio"
	"io"
	"reflect"
	"strconv"
)

type In struct {
	*bufio.Scanner
	buf   []string
	rpos  int
	split bufio.SplitFunc
}

func NewIn(r io.Reader) *In {
	return NewInWithSplitFunc(r, bufio.ScanWords)
}

func NewInWithSplitFunc(r io.Reader, split bufio.SplitFunc) *In {
	m := &In{bufio.NewScanner(r), []string{}, 0, split}
	m.Scanner.Split(m.split)
	return m
}

func (m *In) GetSplitFunc() bufio.SplitFunc {
	return m.split
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

func (m *In) ReadLine() string {
	// note that their type is different: bufio.SplitFunc | func([]uint8, bool) (int, []uint8, error)
	if reflect.ValueOf(m.split).Pointer() != reflect.ValueOf(bufio.ScanLines).Pointer() {
		panic("ReadLine: In.split != bufio.ScanLines")
	}
	return m.ReadString()
}
