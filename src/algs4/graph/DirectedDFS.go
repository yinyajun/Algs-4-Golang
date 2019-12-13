package graph

type DirectedDFS struct {
	marked []bool // marked[v] = true iff v is reachable from source(s)
	count  int    // number of vertices reachable from source(s)
}

//Computes the vertices in digraph G that are reachable from the source vertex s.
func NewDirectedDFS(g *digraph, s int) *DirectedDFS {
	d := &DirectedDFS{}
	d.marked = make([]bool, g.V())
	d.validateVertex(s)
	d.dfs(g, s)
	return d
}

func NewDirectedDFSMultiSources(g *digraph, sources []int) *DirectedDFS {
	d := &DirectedDFS{}
	d.marked = make([]bool, g.V())
	d.validateVertices(sources)
	for _, s := range sources {
		d.dfs(g, s)
	}
	return d
}

func (d *DirectedDFS) validateVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (d *DirectedDFS) validateVertices(sources []int) {
	if len(sources) == 0 {
		panic("validateVertices: empty sources")
	}
	V := len(d.marked)
	for _, s := range sources {
		if s < 0 || s >= V {
			panic("validateVertices: invalid vertex")
		}
	}
}

func (d *DirectedDFS) dfs(g *digraph, v int) {
	d.marked[v] = true
	d.count++
	vAdj := g.Adj(v)
	for hasNext, w := vAdj(); hasNext; hasNext, w = vAdj() {
		if !d.marked[w.(int)] {
			d.dfs(g, w.(int))
		}
	}
}

// Is there a directed path from the source vertex
// (or any of the source vertices) and vertex v?
func (d *DirectedDFS) Marked(v int) bool {
	d.validateVertex(v)
	return d.marked[v]
}

func (d *DirectedDFS) Count() int { return d.count }
