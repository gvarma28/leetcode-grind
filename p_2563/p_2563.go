package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/count-the-number-of-fair-pairs/description/
func main() {
	res := countFairPairs([]int{1,7,9,2,5}, 11, 11)
	fmt.Println(res)
}

// TLE - add binary search
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)

	p, q := 0, n-1

	var res int64 = 0
	for {
		if p >= q {
			break
		}

		var sum int64 = int64(nums[p]) + int64(nums[q])
		if sum >= int64(lower) && sum <= int64(upper) {
			temp := q
			for {
				if p >= temp {
					break
				}
				var sum int64 = int64(nums[p]) + int64(nums[temp])
				if sum >= int64(lower) && sum <= int64(upper) {
					res++
					temp--
				} else {
					break
				}
			}
			p++
		} else if sum < int64(lower) {
			p++
		} else {
			q--
		}
	}

	return res
}
