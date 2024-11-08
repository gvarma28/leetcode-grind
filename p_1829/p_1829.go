package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-xor-for-each-query/description/
func main() {
	res := getMaximumXor([]int{0, 1, 1, 3}, 2)
	fmt.Println(res)
}

func getMaximumXor(nums []int, maximumBit int) []int {
	res := []int{}
	xorVals := []int{}
	curXor := 0
	for _, val := range nums {
		curXor ^= val
		xorVals = append(xorVals, curXor)
	}

	n := len(nums)
	for j := 0; j < n; j++ {
		cur := xorVals[n-j-1]
		max := 0
		for i := 0; i < maximumBit; i++ {
			if (cur>>i)&1 != 1 {
				max += (1 << i)
			}
		}
		res = append(res, max)
	}
	return res
}

// *** TLE - O(N * 2^maximumBit) ***
// func getMaximumXor(nums []int, maximumBit int) []int {
// 	res := []int{}
// 	xorVals := []int{}
// 	curXor := 0
// 	for _, val := range nums {
// 		curXor ^= val
// 		xorVals = append(xorVals, curXor)
// 	}
// 	n := len(nums)
// 	for j := 0; j < n; j++ {
// 		cur := xorVals[n-j-1]
// 		max := 0
// 		maxIdx := 0
// 		for i := 0; i < 1<<maximumBit; i++ {
// 			if cur ^ i > max {
// 				max = cur ^ i
// 				maxIdx = i
// 			}
// 		}
// 		res = append(res, maxIdx)
// 	}
// 	return res
// }
