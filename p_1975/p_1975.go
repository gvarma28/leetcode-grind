package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-matrix-sum/description/
func main() {
	input := [][]int{}
	input = append(input, []int{2, 9, 3})
	input = append(input, []int{5, 4, -4})
	input = append(input, []int{1, 7, 1})
	res := maxMatrixSum(input)
	fmt.Println(res)
}

func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)

	var sum int64 = 0
	min := 1000000
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cur := matrix[i][j]
			if cur <= 0 {
				cnt++
				sum += (int64(cur) * -1)
				if cur*-1 < min {
					min = cur * -1
				}
			} else {
				sum += int64(cur)
				if cur < min {
					min = cur
				}
			}
		}
	}

	if cnt%2 == 1 {
		return sum - 2*int64(min)
	}

	return sum
}
