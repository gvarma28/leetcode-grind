package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/description/
func main() {
	var res bool = digitCount("1210")
	fmt.Println(res)
}

func digitCount(num string) bool {
	count := make(map[string]int)

	for _, val := range num {
		count[string(val)]++
	}

	for i, val := range num {
		digit := int(val - '0')
		if digit != count[strconv.Itoa(i)] {
			return false
		}
	}

	return true
}
