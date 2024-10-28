package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/summary-ranges/description/
func main() {
	res := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	fmt.Println(res)
}

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)

	var str strings.Builder
	cur := -1
	for i := 0; i < n-1; i++ {
		fmt.Println(nums[i], cur)
		if nums[i+1]-nums[i] == 1 {
			if cur == -1 {
				str.WriteString(strconv.Itoa(nums[i]) + "->")
				cur = i;
			}
			continue
		}
		if cur != -1 {
			str.WriteString(strconv.Itoa(nums[i]))
			res = append(res, str.String())
			str.Reset()
		} else {
			temp := strconv.Itoa(nums[i])
			str.WriteString(temp + "->" + temp)
			res = append(res, str.String())
			str.Reset()
		}
		cur = -1
	}

	fmt.Println(cur)


	return res
}
