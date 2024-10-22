package main

import (
	"fmt"
)

// https://leetcode.com/problems/contains-duplicate-ii/description/
func main() {
	res := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	fmt.Println(res)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	mapper := make(map[int]int)

	for i, val := range nums {
		if _, ok := mapper[val]; ok {
			sum := mapper[val] - i
			if sum < 0 {
				sum *= -1
			}
			if sum <= k {
				return true
			}
		}
		mapper[val] = i
	}

	return false
}
