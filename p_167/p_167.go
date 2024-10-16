package main

import (
	"fmt"
)

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func main() {
	res := twoSum([]int{2, 3, 4}, 6)
	fmt.Println(res)
}

func twoSum(numbers []int, target int) []int {

	res := make([]int, 2)
	n := len(numbers)
	p, q := 0, n-1

	for p < q && p < n && q >= 0 {
		sum := numbers[p] + numbers[q]
		if sum == target {
			res[0] = p + 1
			res[1] = q + 1
			break
		}
		if sum < target {
			p++
		} else {
			q--
		}
	}

	return res
}
