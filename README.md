# Algs-4-Golang

`Algs-4-Golang` is Golang port of the Java code in textbook [Algorithms, 4th Edition](https://algs4.cs.princeton.edu/home/)
by Robert Sedgewick and Kevin Wayne. 

It uses `go module` feature, make sure golang version >= 1.13.

## Package Structure
`Algs-4-Golang` has clear structure.
* `abstract`: API(interface) for abstract data structure.  
* `impl`: Implementation(source code & test cases) of abstract data structure. 
* `utils`: Useful tools to support test and debug.

Walk through `impl`, implemented algorithms organized in 6 sub-files corresponding to 6 chapters in Algs4 textbook.

* `impl/fundamentals`
* `impl/sorting`
* `impl/searching`
* `impl/graphs`
* `impl/strings`
* `impl/context`

## Test Cases
Test cases in a `examples` file. Here is a simple usage to test interface `Stack` stated in `abstract/stack.go`

```
go run impl/fundamentals/examples/stack.go LinkedStack < data/tobe.txt
```
## Algorithms Index

#### 1. Fundamentals

| Abstract Data Structure | Implementation                                               |
| ----------------------- | ------------------------------------------------------------ |
| Bag                     | [LinkedBag](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/linked_bag.go)<br />[ResizingArrayBag](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/resizing_array_bag.go) |
| Stack                   | [LinkedStack](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/linked_stack.go)<br />[ResizeArrayStack](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/resizing_array_stack.go) |
| Queue                   | [LinkedQueue](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/linked_queue.go)<br />[ResizingArrayQueue](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/resizing_array_queue.go)                  |
| Union-find              | [QuickFindUF](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_find_uf.go)<br />[QuickUnionUF](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_union_uf.go)<br />[QuickUnionSizeUF](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_union_size_uf.go)<br />[QuickUnionRankUF](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_union_rank_uf.go)<br />[QuickUnionCompressedUF](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_union_compressed_uf.go)<br />[QuickUnionCompressedUF2](https://github.com/yinyajun/Algs-4-Golang/blob/master/src/impl/fundamentals/quick_union_compressed_uf2.go) |



#### 2. Sorting