/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 16:40
 */

package utils

import (
	"runtime"
)

const (
	TrackDepth = 4
)

func Assert(cond bool, info ...string) {
	if !cond {
		defer PrintInTable(assertionErrorStr(info...), TraceTrack())
		panic(StdOut.Sprintf("[Assertion Error]: %s", info))
	}
}

func AssertF(cond bool, format string, a ...interface{}) {
	if !cond {
		info := StdOut.Sprintf(format, a...)
		defer PrintInTable(assertionErrorStr(info), TraceTrack())
		panic(StdOut.Sprintf("[Assertion Error]: %s", info))
	}
}

func assertionErrorStr(info ...string) []string {
	ret := []string{"Assertion Error"}
	for _, i := range info {
		ret = append(ret, indentInfo(i))
	}
	return ret
}

func indentInfo(info string) string {
	return "  " + info
}

func getFrameTrackInfo(skip int) (funcInfo, lineInfo string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	f := runtime.FuncForPC(pc)
	funcInfo = "  " + f.Name()
	lineInfo = StdOut.Sprintf("%s:%d", file, line)
	return
}

// trace track in reverse order (from human perspective)
func TraceTrack() []string {
	infos := []string{}
	for i := 1 + TrackDepth; i >= 2; i-- {
		funcInfo, lineInfo := getFrameTrackInfo(i)
		infos = append(infos, funcInfo)
		infos = append(infos, lineInfo)
	}
	infos = infos[1 : len(infos)-1]
	return infos
}
