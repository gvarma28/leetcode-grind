package main

import (
	"fmt"
)

// problem_url
func main() {
	nums := []int{2, 3, 1, 1, 4}
	res := canJump(nums)
	fmt.Println(res)
}

// valley approach - https://www.youtube.com/watch?v=muDPTDrpS28
func canJump(nums []int) bool {
	max_reachable_idx := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > max_reachable_idx {
			return false
		}
		local_max_reachable := i + nums[i]
		if max_reachable_idx < local_max_reachable {
			max_reachable_idx = local_max_reachable
		}
	}

	return true
}
