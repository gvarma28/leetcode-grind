package main

import (
	"fmt"
)

// https://leetcode.com/problems/contains-duplicate/description/
func main() {
	res := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(res)
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, val := range nums {
		if _, ok := set[val]; ok {
			return true
		}
		set[val] = struct{}{}
	}

	return false
}
