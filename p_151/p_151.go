package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string/
func main() {
	res := reverseWords("  a good   example ")
	fmt.Println(res)
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	var res strings.Builder
	n := len(s)
	start := -1
	for i := n - 1; i >= 0; i-- {
		val := s[i]
		if val == ' ' && start != -1 {
			res.WriteString(s[i+1:start+1] + " ")
			start = -1
            continue
		}
        if val != ' ' && start == -1 {
            start = i
        }
		
	}
	if start != -1 {
		res.WriteString(s[0 : start+1])
	}
	return res.String()
}