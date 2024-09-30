package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-employees-who-met-the-target/description/
func main() {
	var res int = numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2)
	// var res int = numberOfEmployeesWhoMetTarget([]int{5, 1, 4, 2, 2}, 6)
	fmt.Println(res)
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for _, v := range hours {
		if v >= target {
			res++
		}
	}
	return res
}
