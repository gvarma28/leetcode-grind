package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/count-the-number-of-fair-pairs/description/
func main() {
	res := countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)
	fmt.Println(res)
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)

	var res int64 = 0
	for i := 0; i < n; i++ {
		cur := nums[i]

		max := upper - cur
		min := lower - cur

		upperIdx := bs(nums, i+1, n-1, max+1)
		lowerIdx := bs(nums, i+1, n-1, min)

		res += int64(upperIdx) - int64(lowerIdx)
	}

	return res
}

func bs(nums []int, l int, r int, target int) int {
	for {
		if l > r {
			break
		}

		var mid int = (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return r
}
