package main

import (
	"fmt"
)

// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/
func main() {
	res := largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
	fmt.Println(res)
}

func largestCombination(candidates []int) int {
	max := 0
	for _, val := range candidates {
		if val > max {
			max = val
		}
	}

	res := 0
	for bit := 1; bit <= max; bit <<= 1 {
		count := 0
		for _, val := range candidates {
			if val&bit > 0 {
				count++
			}
		}
		if count > res {
			res = count
		}
	}

	return res
}

// func largestCombination(candidates []int) int {

// 	setBitCounter := [32]int{}

// 	for _, val := range candidates {
// 		for i := 0; i < 32; i++ {
// 			if (val>>i)%2 == 1 {
// 				setBitCounter[i]++
// 			}
// 		}
// 	}

// 	res := 0
// 	for i := 0; i < 32; i++ {
// 		if setBitCounter[i] > res {
// 			res = setBitCounter[i]
// 		}
// 	}

// 	return res
// }
