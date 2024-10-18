package main

import (
	"fmt"
)

// https://leetcode.com/problems/ransom-note/description/
func main() {
	res := canConstruct("aa", "aab")
	fmt.Println(res)
}

func canConstruct(ransomNote string, magazine string) bool {

	count := make([]int, 26)

	for _, val := range magazine {
		count[val - 'a']++
	}

	for _, val := range ransomNote {
		if count[val - 'a'] == 0 {
			return false
		}
		count[val - 'a']--
	}

	return true;
}