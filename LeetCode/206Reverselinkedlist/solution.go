package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	result := &ListNode{}
	n := len(arr)
	list := result
	for i := range arr {
		list.Val = arr[n-1-i]
		if i != n-1 {
			list.Next = &ListNode{}
			list = list.Next
		}
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rest = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr = &ListNode{Val: head.Val}
	if head.Next != nil {
		head.Val = head.Next.Val
		head.Next = current
	}

}

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	//l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}

	l2 := reverseList(&l1)
	for l2 != nil {
		fmt.Println("Val", l2.Val)
		l2 = l2.Next
	}
	//fmt.Println()
	//fmt.Println("rev", l2.Val)
	//efmt.Println("adsfdf", printRev(&l1))
}
