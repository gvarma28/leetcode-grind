package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-1-bits/description/
func main() {
	res := hammingWeight(2147483645)
	fmt.Println(res)
}

func hammingWeight(n int) int {
	res := 0
	for i := 0; i < 32; i++ {
		if (n>>i)%2 == 1 {
			res++
		}
	}
	return res
}
