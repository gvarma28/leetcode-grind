package main

import (
	"fmt"
)

// https://leetcode.com/problems/rotate-array/description/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = (k % n)
	if k == 0 || n == 1 {
		return
	}
	reverse(&nums, 0, n-k-1)
	reverse(&nums, n-k, n-1)
	reverse(&nums, 0, n-1)
}

func reverse(nums *[]int, start int, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}

/* brute force O(n*k)
func rotate(nums []int, k int) {

	for i := 0; i < k; i++ {
		nums = rotateOnce(nums)
	}
}

func rotateOnce(nums []int) []int {

	cur := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := nums[(i+1)%n]
		nums[(i+1)%n] = cur;
		cur = temp
	}

	return nums
*/
