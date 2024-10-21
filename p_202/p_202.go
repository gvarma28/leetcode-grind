package main

import (
	"fmt"
)

// https://leetcode.com/problems/happy-number/description/
func main() {
	res := isHappy(19)
	// res := isHappy(2)
	fmt.Println(res)
}

func isHappy(n int) bool {

	squareMap := make(map[int]struct{})

	for n > 0 {
		if n == 1 {
			return true
		}
		if _, ok := squareMap[n]; ok {
			return false
		} else {
			squareMap[n] = struct{}{}
		}
		tempN := 0
		for n > 0 {
			lastDigit := n % 10
			tempN += lastDigit * lastDigit
			n = n / 10
		}
		n = tempN
	}

	return true
}
