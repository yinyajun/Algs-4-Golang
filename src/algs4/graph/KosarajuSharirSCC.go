package graph

/**
* determining the strong components in a Digraph
* This implementation uses the Kosaraju-Sharir algorithm
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type KosarajuSharirSCC struct {
	marked []bool
	id     []int
	count  int
}

func NewKosarajuSharirScc(g *Digraph) *KosarajuSharirSCC {
	scc := &KosarajuSharirSCC{}

	// compute reverse postorder of reverse Graph
	dfs := NewDepthFirstOrder(g.Reverse())

	// run DFS on G, using reverse postorder to guide calculation
	scc.marked = make([]bool, g.V())
	scc.id = make([]int, g.V())
	order := dfs.ReversePostOrder()
	for v := order.Next(); v != nil; v = order.Next() {
		if !scc.marked[v.(int)] {
			scc.dfs(g, v.(int))
			scc.count++
		}
	}
	// check that id[] gives strong components
	// todo: check

	return scc
}

func (scc *KosarajuSharirSCC) dfs(g *Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count

	vAdj := g.Adj(v)
	for w := vAdj.Next(); w != nil; w = vAdj.Next() {
		if !scc.marked[w.(int)] {
			scc.dfs(g, w.(int))
		}
	}
}

func (scc *KosarajuSharirSCC) Count() int { return scc.count }

func (scc *KosarajuSharirSCC) Id(v int) int {
	scc.validateVertex(v)
	return scc.id[v]
}

func (scc *KosarajuSharirSCC) validateVertex(v int) {
	V := len(scc.marked)
	if v < 0 || v >= V {
		panic("validateVertex: invalid vertex")
	}
}

func (scc *KosarajuSharirSCC) Connected(v, w int) bool { return scc.Id(v) == scc.Id(w) }

func (scc *KosarajuSharirSCC) check(g *Digraph) bool {
	tc := NewTransitiveClosure(g)
	for v := 0; v < g.V(); v++ {
		for w := 0; w < g.V(); w++ {
			if scc.Connected(v, w) != tc.Reachable(v, w) && tc.Reachable(w, v) {
				return false
			}
		}
	}
	return true
}
