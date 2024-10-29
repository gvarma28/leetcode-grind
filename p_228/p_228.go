package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/summary-ranges/description/
func main() {
	res := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	fmt.Println(res)
}

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)

	if n == 0 {
		return res
	}

	var str strings.Builder
	firstElem := true 
	for i := 0; i < n; i++ {
		if i < n-1 && nums[i+1]-nums[i] == 1 {
			if firstElem {
				str.WriteString(strconv.Itoa(nums[i]) + "->")
				firstElem = false
			}
			continue
		}
		str.WriteString(strconv.Itoa(nums[i]))
		res = append(res, str.String())
		str.Reset()
		firstElem = true
	}

	return res
}
