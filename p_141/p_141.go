package main

import (
	"fmt"
)

// problem_url
func main() {

	values := []int{3,2,0,-4}
	pos := 1

	head := ListNode{
		Val: 3,
	}

	curLink := &head
	var posLink *ListNode
	for i, val := range values {
		if i == 0 {
			if pos == 0 {
				posLink = &head
			}
			continue
		}
		curNode := ListNode{
			Val: val,
		}
		curLink.Next = &curNode
		curLink = &curNode
		if i == pos {
			posLink = &curNode
		}
	}
	fmt.Println(posLink)
	if pos >= 0 {
		if pos != 0 {
			curLink.Next = posLink.Next
		} else {
			curLink.Next = posLink
		}
	}

	temp := head
	for i := 0; i < len(values); i++ {
		fmt.Println(temp, temp.Next)
		if temp.Next != nil {
			temp = *temp.Next
		}
	}

	fmt.Println(hasCycle(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: cleanup and complete this
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	var p *ListNode = head
	var q *ListNode = head

	for p.Next != nil && q.Next != nil && p != q {
		fmt.Println(p, q)
		if p.Next.Next != nil {
			p = p.Next.Next
		} else {
			return false
		}
		q = q.Next
	}

	fmt.Println(p, q)

	if p.Next == nil || q.Next == nil || p == q {
		return false
	}

	return true
}
