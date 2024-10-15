package main

import (
	"fmt"
	"unicode/utf8"
)

// https://leetcode.com/problems/is-subsequence/description/
func main() {
	res := isSubsequence("axc", "ahbgdc")
	fmt.Println(res)
}

func isSubsequence(s string, t string) bool {
	n, m := utf8.RuneCountInString(s), utf8.RuneCountInString(t)
	p, q := 0, 0

	for p < n && q < m {
		if s[p] == t[q] {
			p++
		}
		q++
	}

	if p == n {
		return true
	}
	return false
}
