package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-if-array-can-be-sorted/description/
func main() {
	res := canSortArray([]int{3, 16, 8, 4, 2})
	fmt.Println(res)
}

func canSortArray(nums []int) bool {
	n := len(nums)
	bits := []int{}
	for i := 0; i < n; i++ {
		bits = append(bits, countSetBits(nums[i]))
	}

	p, q := 0, 0
	for {
		if p >= n || q >= n {
			break
		}
		if p == q {
			q++
			continue
		}

		if nums[p] > nums[q] && bits[p] == bits[q] {
			nums[p], nums[q] = nums[q], nums[p]
			bits[p], bits[q] = bits[q], bits[p]
			p = 0
			q = 0
		} else {
			p++
			q++
		}
	}

	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func countSetBits(n int) int {
	count := 0
	for i := 0; i <= 8; i++ {
		if (n>>i)%2 == 1 {
			count++
		}
	}
	return count
}
