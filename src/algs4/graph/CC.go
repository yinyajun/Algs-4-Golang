package graph

/**
* determining the connected components in an undirected graph.
* his implementation uses depth-first search.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type cc struct {
	marked []bool
	count  int   // number of connected components
	id     []int //id[v] = id of connected component containing v
	size   []int //size[id] = number of vertices in given component
}

func NewCC(g *graph) *cc {
	c := &cc{}
	c.marked = make([]bool, g.V())
	c.id = make([]int, g.V())
	c.size = make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if !c.marked[v] {
			c.dfs(g, v)
			c.count++
		}
	}
	return c
}

//todo: edge weight

func (c *cc) dfs(g *graph, v int) {
	c.marked[v] = true
	c.id[v] = c.count
	c.size[c.count]++
	vAdj := g.Adj(v)
	for w := vAdj.Next(); w != nil; w = vAdj.Next() {
		if !c.marked[w.(int)] {
			c.dfs(g, w.(int))
		}
	}
}

func (c *cc) Id(v int) int {
	c.validateVertex(v)
	return c.id[v]
}

func (c *cc) Count() int     { return c.count }
func (c *cc) Size(v int) int { return c.size[c.Id(v)] }

func (c *cc) Connected(v, w int) bool {
	c.validateVertex(v)
	c.validateVertex(w)
	return c.Id(v) == c.Id(w)

}

func (c *cc) validateVertex(v int) {
	V := len(c.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
