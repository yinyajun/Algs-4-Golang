# Algs-4-Golang
Golang translations of Robert Sedgewick's Java Algorthms. Details see [http://algs4.cs.princeton.edu].

## Package Structure
* `data/`: data for test cases, full data should be download from [http://algs4.cs.princeton.edu]
* `src/algs4/`: source code of data structure and algorithm 
* `src/test/`: test cases 
* `src/util/`: utility tools, e.g., Scanner, Iterator, etc.


## Usage
1. Set `GOPATH` properly to ensure this package can be find. 
2. Run test cases. Each test case correspond to an implementation in `src/algs4`.
    E.g. `go run src/test/graph.go < data/tinyG.txt`

    ```
    13 vertices, 13 edges 
    0:6 2 1 5 
    1:0 
    2:0 
    3:5 4 
    4:5 6 3 
    5:3 4 0 
    6:0 4 
    7:8 
    8:7 
    9:11 10 12 
    10:9 
    11:9 12 
    12:11 9
    ```
