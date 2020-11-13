package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var root ListNode
	tail := &root
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			tail.Next = &ListNode{Val: l1.Val, Next: nil}
			l1 = l1.Next
		} else {
			if l2 != nil {
				tail.Next = &ListNode{Val: l2.Val, Next: nil}
				l2 = l2.Next
			}
		}
		tail = tail.Next
	}
	return root.Next
}

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println("result", mergeTwoLists(&l1, &l2))
}
