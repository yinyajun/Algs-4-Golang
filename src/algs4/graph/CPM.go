package graph

import (
	"util"
	"fmt"
)

/**
* provides a client that solves the
*  parallel precedence-constrained job scheduling problem
*  via the critical path method. It reduces the problem
*  to the longest-paths problem in edge-weighted DAGs.
*  It builds an edge-weighted digraph (which must be a DAG)
*  from the job-scheduling problem specification,
*  finds the longest-paths tree, and computes the longest-paths
*  lengths (which are precisely the start times for each job).
*
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

type CPM struct{}

func NewCPM(in *util.In) {
	// number of jobs
	n := in.ReadInt()

	// source and sink
	source := 2 * n
	sink := 2*n + 1

	// build network
	g := NewEdgeWeightedDigraph(2*n + 2)
	for i := 0; i < n; i++ {
		duration := in.ReadFloat()

		g.AddEdge(NewDirectedEdge(i, i+n, duration))
		g.AddEdge(NewDirectedEdge(source, i, 0.0))
		g.AddEdge(NewDirectedEdge(i+n, sink, 0.0))

		m := in.ReadInt()
		for j := 0; j < m; j++ {
			successor := in.ReadInt()
			g.AddEdge(NewDirectedEdge(i+n, successor, 0.0))
		}
	}

	// compute longest path
	lp := NewAcyclicLP(g, source)

	// print results
	fmt.Println(" job   start  finish")
	fmt.Println("--------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%4d %7.1f %7.1f\n", i, lp.DistTo(i), lp.DistTo(i+n))
	}
	fmt.Printf("Finish time: %7.1f\n", lp.DistTo(sink))
}
