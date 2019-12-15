package graph

import (
	"bufio"
	"os"
	"strings"
	"util"
)

/**
* symbol Digraph
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type SymbolDigraph struct {
	st    map[string]int
	keys  []string
	graph *Digraph
}

func NewSymbolDigraph(filename string, delimiter string) *SymbolDigraph {
	sd := &SymbolDigraph{}

	// First pass builds the index by reading strings to associate
	// distinct strings with an index
	sd.st = make(map[string]int)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	in := util.NewInWithSplitFunc(f, bufio.ScanLines)
	for in.HasNext() {
		a := strings.Split(in.ReadLine(), delimiter)
		for i := 0; i < len(a); i++ {
			if _, ok := sd.st[a[i]]; !ok {
				sd.st[a[i]] = len(sd.st)
			}
		}
	}

	// inverted index to get string keys in an array
	sd.keys = make([]string, len(sd.st))
	for name, idx := range sd.st {
		sd.keys[idx] = name
	}
	f.Close()

	// second pass builds the Digraph by connecting first vertex on each
	// line to all others
	sd.graph = NewDigraph(len(sd.st))
	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	in = util.NewInWithSplitFunc(f, bufio.ScanLines)
	for in.HasNext() {
		a := strings.Split(in.ReadLine(), delimiter)
		v := sd.st[a[0]]
		for i := 1; i < len(a); i++ {
			w := sd.st[a[i]]
			sd.graph.AddEdge(v, w)
		}
	}
	f.Close()
	return sd
}

func (sd *SymbolDigraph) Contains(key string) bool {
	_, ok := sd.st[key]
	return ok
}

func (sd *SymbolDigraph) Index(key string) int {
	return sd.st[key]
}

func (sd *SymbolDigraph) Name(v int) string {
	sd.validateVertex(v)
	return sd.keys[v]
}

func (sd *SymbolDigraph) Digraph() *Digraph { return sd.graph }

func (sd *SymbolDigraph) validateVertex(v int) {
	V := sd.graph.V()
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
