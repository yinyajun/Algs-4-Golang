package graph

/**
* determining the connected components in an undirected Graph.
* his implementation uses depth-first search.
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type CC struct {
	marked []bool
	count  int   // number of connected components
	id     []int //id[v] = id of connected component containing v
	size   []int //size[id] = number of vertices in given component
}

func NewCC(g *Graph) *CC {
	c := &CC{}
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

//todo: Edge weight

func (c *CC) dfs(g *Graph, v int) {
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

func (c *CC) Id(v int) int {
	c.validateVertex(v)
	return c.id[v]
}

func (c *CC) Count() int     { return c.count }
func (c *CC) Size(v int) int { return c.size[c.Id(v)] }

func (c *CC) Connected(v, w int) bool { return c.Id(v) == c.Id(w) }

func (c *CC) validateVertex(v int) {
	V := len(c.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}
