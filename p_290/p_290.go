package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/word-pattern/description/
func main() {
	res := wordPattern("abba", "dog dog dog dog")
	fmt.Println(res)
}

func wordPattern(pattern string, s string) bool {
	// words := strings.Fields(s) // splits around whitespace
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	dict := make(map[rune]string)
	dict2 := make(map[string]rune)

	for i, v := range pattern {
		if _, ok := dict[v]; ok {
			if dict[v] != words[i] {
				return false
			}
			continue
		}

		if _, ok := dict2[words[i]]; ok {
			if dict2[words[i]] != v {
				return false
			}
			continue
		}

		dict[v] = words[i]
		dict2[words[i]] = v
	}

	return true
}
