package main

import (
	"fmt"
)

// https://leetcode.com/problems/symmetric-tree/
func main() {
	fmt.Println("ThankYouForReadingMyCode")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return rec(root.Left, root.Right)
}

func rec(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		if left == nil && right == nil {
			return true
		}
		return false
	}

	if left.Val != right.Val {
		return false
	}

	rec1 := rec(left.Left, right.Right)
	rec2 := rec(left.Right, right.Left)

	if !rec1 || !rec2 {
		return false
	}
	return true
}
