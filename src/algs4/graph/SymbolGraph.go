package graph

import (
	"os"
	"util"
	"bufio"
)

/**
* symbol graph
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type symbolGraph struct {
	st    map[string]int // string -> index
	keys  []string       // index  -> string
	graph *graph         // the underlying graph
}

func NewSymbolGraph(filename string, delimiter string) *symbolGraph {
	sg := &symbolGraph{}
	sg.st = make(map[string]int)
	f, _ := os.Open(filename)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	in := util.NewInWithSplitFunc(f, bufio.ScanLines)


	return sg

}
