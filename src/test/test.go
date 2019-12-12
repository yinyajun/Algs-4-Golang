package main

import (
	"fmt"
	"strings"
)

func main() {
	//d, err := os.Getwd()
	//fmt.Println(d)
	//f, err := os.Open(`D:\work\git_repo\Algs-4-Golang\src\test\graph.go`)
	//fmt.Println(err)
	//defer func() {
	//	if f != nil {
	//		f.Close()
	//	}
	//}()
	////contentByte,err:=ioutil.ReadAll(f)
	////fmt.Println(string(contentByte))
	//c := bufio.ScanLines
	//in := util.NewInWithSplitFunc(f, c)
	//fmt.Println(in.ReadLine())
	s := strings.Split("abc,abc", "|")
	fmt.Println(s, len(s))
}
