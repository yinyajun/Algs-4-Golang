/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 20:08
 */

package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	StdOut *stdOut
)

func init() {
	StdOut = NewStdOut(os.Stdout)
}

type stdOut struct {
	*bufio.Writer
}

func NewStdOut(w io.Writer) *stdOut {
	writer := bufio.NewWriter(w)
	return &stdOut{writer}
}

func (o *stdOut) Println(a ...interface{}) {
	defer o.Flush()
	fmt.Fprintln(o, a...)
}

func (o *stdOut) Print(a ...interface{}) {
	defer o.Flush()
	fmt.Fprint(o, a...)
}

func (o *stdOut) Printf(format string, a ...interface{}) {
	defer o.Flush()
	fmt.Fprintf(o, format, a...)
}

func (o *stdOut) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
