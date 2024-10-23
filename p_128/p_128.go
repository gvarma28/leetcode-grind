package main

import (
	"fmt"
)

// problem_url
func main() {
	res := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, val := range nums {
		set[val] = true
	}

	res := 0
	for key := range set {
		if !set[key-1] {
			count := 1
			nextKey := key + 1

			for set[nextKey] {
				count++
				nextKey++
			}

			if count > res {
				res = count
			}
		}
	}

	return res

}
