package main

import (
	"fmt"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
func main() {
	res := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	fmt.Println(res)
}

// CFOT
func removeDuplicates(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		j := i
		for ; j < n && nums[j] == nums[i]; j++ {
			if j-i >= 2 {
				nums[j] = 100000
			}
		}
		i = j - 1
	}

	counter := 0
	for i := 0; i < n-counter; i++ {
		if nums[i] == 100000 {
			for j := i; j < n-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[n-1] = 100000
			counter++
			i--
		}
	}

	return n - counter
}
