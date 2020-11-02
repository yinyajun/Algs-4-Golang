/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/11/1 13:45
 */

package utils

import "path"

const (
	FileNameTruncateLength = 4
)

func Split(name string) (string, string) {
	dirName, fileName := path.Split(name)
	if len(dirName) > 0 {
		dirName = dirName[:len(dirName)-1]
	}
	return dirName, fileName
}

// Truncate filePath which is too long.
// level is fileName length
// return: (dirName, fileName)
// e.g. TruncateSplit("/a/b/c/d/e", 3) = (/a/b, c/d/e)
func TruncateSplit(filePath string, level int) (string, string) {
	if level == 0 {
		return filePath, ""
	}
	dirName, fileName := TruncateSplit(filePath, level-1)
	a, b := Split(dirName)
	c := path.Join(b, fileName)
	return a, c
}

func TruncatePath(filePath string, level int) string {
	_, path := TruncateSplit(filePath, level)
	return path
}
