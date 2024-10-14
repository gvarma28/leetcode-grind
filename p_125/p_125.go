package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// https://leetcode.com/problems/valid-palindrome/description/
func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := utf8.RuneCountInString(s)

	p, q := 0, n-1

	for p < q {
		rp, _ := utf8.DecodeRuneInString(s[p:])
		rq, _ := utf8.DecodeLastRuneInString(s[:q+1])

		if !unicode.IsLetter(rp) && !unicode.IsDigit(rp) {
			p++
			continue
		}
		if !unicode.IsLetter(rq) && !unicode.IsDigit(rq) {
			q--
			continue
		}

		if rp != rq {
			return false
		}

		p++
		q--

	}
	return true
}
