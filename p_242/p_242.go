package main

import (
	"fmt"
)

// https://leetcode.com/problems/valid-anagram/description/
func main() {
	res := isAnagram("anagram", "nagaram")
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	countArr := make([]int, 26)
	for _, val := range s {
		countArr[val-'a']++
	}
	for _, val := range t {
		countArr[val-'a']--
	}
	for _, val := range countArr {
		if val != 0 {
			return false
		}
	}
	return true

}
