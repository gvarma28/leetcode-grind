package main

import (
	"fmt"
)

// https://leetcode.com/problems/single-number/description/
func main() {
	res := singleNumber([]int{2, 2, 4})
	fmt.Println(res)
}

func singleNumber(nums []int) int {

	res := 0
	for _, val := range nums {
		res ^= val
	}

	return res
}
