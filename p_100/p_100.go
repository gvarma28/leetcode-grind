package main

import (
	"fmt"
)

// https://leetcode.com/problems/same-tree/
func main() {
	fmt.Println("ThankYouForReadingMyCode")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return rec(p, q)
}

func rec(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := rec(p.Left, q.Left)
	right := rec(p.Right, q.Right)
	if !left || !right {
		return false
	}
	return true
}
