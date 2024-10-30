package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func main() {
	fmt.Println("ThankYouForReadingMyCode")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return rec(root, 0)
}

func rec(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	left := rec(root.Left, count+1)
	right := rec(root.Right, count+1)
	if left > right {
		return left
	}
	return right
}
