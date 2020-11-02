/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 12:40
 */

package utils

import "flag"

var (
	Flag *argParser
	Arg0 string
	Arg1 string
	Arg2 string
	Arg3 string
	Arg4 string
	Arg5 string
	Arg6 string
)

func init() {
	Flag = &argParser{}
}

type argParser struct {
	index int
	value string
}

func (a *argParser) Arg(i int, value ...string) string {
	flag.Parse()
	ret := flag.Arg(i)
	if ret != "" {
		return ret
	}
	if len(value) == 0 {
		return ""
	}
	return value[0]
}
