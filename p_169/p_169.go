package main

import (
	"fmt"
)

// https://leetcode.com/problems/majority-element/description/
func main() {
	var res int = majorityElement([]int{3, 2, 3})
	fmt.Println(res)
}

func majorityElement(nums []int) int {
	count := make(map[int]int)
	n := len(nums)
	for _, val := range nums {
		count[val]++
		if count[val] > n/2 {
			return val
		}
	}
	return 0
}

// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)/2]
// }
