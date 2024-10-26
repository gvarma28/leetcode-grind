package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/description/
func main() {
	res := removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"})
	fmt.Println(res)
}

func removeSubfolders(folder []string) []string {
	hashTable := make(map[string]struct{})
	for _, val := range folder {
		hashTable[val] = struct{}{}
	}

	res := []string{}
out:
	for _, val := range folder {
		dirArr := strings.Split(strings.Trim(val, "/"), "/")

		var curPath strings.Builder
		dirArrLen := len(dirArr)
		for i := 0; i < dirArrLen-1; i++ {
			dir := dirArr[i]
			if dir == "" {
				continue
			}
			curPath.WriteString(("/" + dir))
			if _, ok := hashTable[curPath.String()]; ok {
				continue out
			}
		}

		res = append(res, val)
	}

	return res
}
