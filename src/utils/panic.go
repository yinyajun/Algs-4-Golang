/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/2 11:58
 */

package utils

func Panic(info ...string) {
	defer PrintInTable(panicStr(info...), TraceTrack())
	panic(StdOut.Sprintf("[Panic]: %s", info))
}

func PanicF(format string, a ...interface{}) {
	info := StdOut.Sprintf(format, a...)
	defer PrintInTable(panicStr("  "+info), TraceTrack())
	panic(StdOut.Sprintf("[Panic]: %s", info))
}

func panicStr(info ...string) []string {
	ret := []string{"Panic"}
	for _, i := range info {
		ret = append(ret, indentInfo(i))
	}
	return ret
}
