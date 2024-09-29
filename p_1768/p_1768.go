package main

import "fmt"
import "unicode/utf8"
import "strings"

// https://leetcode.com/problems/merge-strings-alternately/description/
func main() {
	var res string = mergeAlternately("abc", "pqr")
	// var res string = mergeAlternately("ab", "pqrs")
	// var res string = mergeAlternately("abcd", "pq")
	fmt.Println(res)
}

func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	n := utf8.RuneCountInString(word1)
	m := utf8.RuneCountInString(word2)

	i := 0
	for ; i < n && i < m; i++ {
		res.WriteString(string(word1[i]))
		res.WriteString(string(word2[i]))
	}
	for ; i < n; i++ {
		res.WriteString(string(word1[i]))
	}
	for ; i < m; i++ {
		res.WriteString(string(word2[i]))
	}
	return res.String()
}
