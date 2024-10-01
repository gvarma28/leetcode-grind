package main

import (
	"fmt"
)

// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/description/
func main() {
	res := getSneakyNumbers([]int{0, 1, 1, 0})
	// res := getSneakyNumbers([]int{0, 3, 2, 1, 3, 2})
	// res := getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2})
	fmt.Println(res)
}

func getSneakyNumbers(nums []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 2)

	first := true
	for _, val := range nums {
		if _, ok := set[val]; ok {
			if first {
				res[0] = val
				first = false
			} else {
				res[1] = val
			}
		}
		set[val] = struct{}{}
	}
	return res
}
