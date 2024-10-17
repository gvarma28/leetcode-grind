package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-size-subarray-sum/description/
func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	// res := minSubArrayLen(4, []int{1, 4, 4})
	// res := minSubArrayLen(11, []int{1,1,1,1,1,1,1,1})
	fmt.Println(res)
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 1 && nums[0] >= target {
		return 1
	}

	p, q := 0, 0
	res := 1000000
	sum := 0
	for p < n && q < n {
		fmt.Println(p, q)
		if sum+nums[q] >= target {
			if res > q-p {
				res = q - p + 1
			}
			sum -= nums[p]
			if p == q {
				q++
			}
			p++
		} else {
			sum += nums[q]
			q++
		}
	}

	if res == 1000000 {
		return 0
	}
	return res
}
