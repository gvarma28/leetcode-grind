package main

import (
	"fmt"
)

// https://leetcode.com/problems/subsets/description/
func main() {
	res := subsets([]int{1,2,3})

	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	for i := 0; i < 1<<n; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			if (i>>j)%2 == 1 {
				temp = append(temp, nums[j])
			}
		}
		res = append(res, temp)
	}
	return res
}
