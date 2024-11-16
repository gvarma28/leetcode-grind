package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/
func main() {
	res := resultsArray([]int{1, 4}, 1)
	fmt.Println(res)
}

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	if n == 1 && k == 1 {
		return nums
	}

	res := []int{}
	p, q := 0, 0
	validCounter := 0
	for {
		if q >= n {
			break
		}
		if q > 0 && nums[q] == nums[q-1]+1 {
			validCounter++
		}

		if q-p+1 == k {
			if validCounter+1 == k {
				res = append(res, nums[q])
			} else {
				res = append(res, -1)
			}

			if p < n-1 && nums[p]+1 == nums[p+1] {
				validCounter--
			}
			p++
		}
		q++
	}

	return res
}
