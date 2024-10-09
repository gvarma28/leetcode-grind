package main

import (
	"fmt"
)

// problem_url
func main() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	res := 0
	cur := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] - cur > 0 {
			res += prices[i] - cur
		}
		cur = prices[i]
	}
	return res;
}
