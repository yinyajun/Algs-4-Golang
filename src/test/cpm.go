package main

import (
	"os"
	"algs4/graph"
	. "util"
)

/**
* $ go run src/test/cpm.go < data/jobsPC.txt
*  job   start  finish
* --------------------
*    0     0.0    41.0
*    1    41.0    92.0
*    2   123.0   173.0
*    3    91.0   127.0
*    4    70.0   108.0
*    5     0.0    45.0
*    6    70.0    91.0
*    7    41.0    73.0
*    8    91.0   123.0
*    9    41.0    70.0
* Finish time:   173.0
* @see
* @author Golang translation by Yajun Yin from Java by Robert Sedgewick and Kevin Wayne.
 */

func main() {
	in := NewIn(os.Stdin)
	graph.NewCPM(in)
}
