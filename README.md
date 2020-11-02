# Algs-4-Golang

`Algs-4-Golang` is Golang port of the Java code in textbook [Algorithms, 4th Edition](https://algs4.cs.princeton.edu/home/)
by Robert Sedgewick and Kevin Wayne. 

## Package Structure
`Algs-4-Golang` has clear structure.
* `src/abstract`: API(interface) for abstract data structure.  
* `src/impl`: Implementation(source code) of abstract data structure. 
* `src/utils`: Useful tools to support test and debug.
* `download.sh`: Script for downloading test data files 
* `run.sh`：Script for setting `GOPATH` and wrapping `go run` command, you can use this to replace `go run`

Walk through `src/impl`, implemented algorithms organized in 6 sub-files corresponding to 6 chapters in Algs4 textbook.

* `src/impl/fundamentals`
* `src/impl/sorting`
* `src/impl/searching`
* `src/impl/graphs`
* `src/impl/strings`
* `src/impl/context`

## Test Cases
In each sub-file under `src/impl`, there are test cases in a `examples` file.
To avoid setting `GOPATH` environment variable, you could use `run.sh` instead of `go run` command.

Here is a simple usage to test interface `Stack` stated in `src/abstract/stack.go`
```
./run.sh src/impl/fundamentals/examples/stack.go LinkedStack < data/tobe.txt
```
Explanations:
* `./run.sh`: Run command, wrapper of `go run` but setting proper `GOPATH`.
* `src/impl/fundamentals/examples/stack.go`: Test case file, for testing abstract data structure.  
* `LinkedStack`: Arg, for specifying a concrete implementation.
* `< data/tobe.txt`： Input redirect.

## Algorithms Index

#### 1. Fundamentals

| Abstract Data Structure | Implementation                                               |
| ----------------------- | ------------------------------------------------------------ |
| Bag                     | LinkedBag<br />ResizingArrayBag                              |
| Stack                   | LinkedStack<br />ResizeArrayStack                            |
| Queue                   | LinkedQueue<br />ResizingArrayQueue                          |
| Union-find              | QuickFindUF<br />QuickUnionUF<br />QuickUnionSizeUF<br />QuickUnionRankUF<br />QuickUnionCompressedUF<br />QuickUnionCompressedUF2 |



#### 2. Sorting