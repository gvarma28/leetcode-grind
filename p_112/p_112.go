package main

import (
	"fmt"
)

// https://leetcode.com/problems/path-sum/
func main() {
	fmt.Println("ThankYouForReadingMyCode")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	var left bool
	if root.Left != nil {
		left = hasPathSum(root.Left, targetSum-root.Val)
	}

	var right bool
	if root.Right != nil {
		right = hasPathSum(root.Right, targetSum-root.Val)
	}

	return left || right
}