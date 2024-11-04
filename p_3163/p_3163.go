package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/string-compression-iii/description/
func main() {
	res := compressedString("abbcde")
	fmt.Println(res)
}

func compressedString(word string) string {
	var res strings.Builder
	n := len(word)
	for i := 0; i < n; i++ {
		count := 1
		cur := word[i]
		for {
			if i < n-1 && cur == word[i+1] && count < 9 {
				i++
				count++
			} else {
				break
			}
		}
		res.WriteString(strconv.Itoa(count) + string(cur))
	}
	return res.String()
}
