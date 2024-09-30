package main

import (
	"fmt"
	"unicode/utf8"
)

// https://leetcode.com/problems/score-of-a-string/description/
func main() {
	var res int = scoreOfString("hello")
	// var res int = scoreOfString("zaz")
	fmt.Println(res)
}

func scoreOfString(s string) int {
	n, res := utf8.RuneCountInString(s), 0
	for i, v := range s {
		if i+1 == n {
			break
		}
		cur := int(v) - int(s[i+1])
		if cur < 0 {
			cur = -1 * cur
		}
		res += cur
	}
	return res
}
