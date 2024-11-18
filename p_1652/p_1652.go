package main

import (
	"fmt"
)

// https://leetcode.com/problems/defuse-the-bomb/description/
func main() {
	res := decrypt([]int{2, 4, 9, 3}, -2)
	fmt.Println(res)
}

func decrypt(code []int, k int) []int {
	res := []int{}
	n := len(code)
	neg := false
	if k == 0 {
		for i:=0; i<n; i++ {
			res = append(res, 0)
		}
		return res
	}
	if k < 0 {
		neg = true
		k = -1 * k
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			if !neg {
				sum += code[(i+1+j)%n]
			} else {
				idx := n + i - 1 - j
				if idx < 0 {
					idx = -1 * idx
				}
				sum += code[idx%n]
			}
		}
		res = append(res, sum)
	}

	return res
}
