package graph

import (
	"bufio"
	"os"
	"strings"
	"util"
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

	// First pass builds the index by reading strings to associate
	// distinct strings with an index
	sg.st = make(map[string]int)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	in := util.NewInWithSplitFunc(f, bufio.ScanLines)
	for in.HasNext() {
		a := strings.Split(in.ReadLine(), delimiter)
		for i := 0; i < len(a); i++ {
			if _, ok := sg.st[a[i]]; !ok {
				sg.st[a[i]] = len(sg.st)
			}
		}
	}

	// inverted index to get string keys in an array
	sg.keys = make([]string, len(sg.st))
	for name, idx := range sg.st {
		sg.keys[idx] = name
	}
	f.Close()

	// second pass builds the graph by connecting first vertex on each
	// line to all others
	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	in = util.NewInWithSplitFunc(f, bufio.ScanLines)
	for in.HasNext() {
		a := strings.Split(in.ReadLine(), delimiter)
		v := sg.st[a[0]]
		for i := 1; i < len(a); i++ {
			w := sg.st[a[i]]
			sg.graph.AddEdge(v, w)
		}
	}
	f.Close()
	return sg
}

func (sg *symbolGraph) Contains(key string) bool {
	_, ok := sg.st[key]
	return ok
}

func (sg *symbolGraph) Index(key string) int {
	return sg.st[key]
}

func (sg *symbolGraph) Name(v int) string {
	sg.validateVertex(v)
	return sg.keys[v]
}

func (sg *symbolGraph) Graph() *graph { return sg.graph }

func (sg *symbolGraph) validateVertex(v int) {
	V := sg.graph.V()
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
