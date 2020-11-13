package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var root ListNode
	tail := &root
	for head != nil {
		if head.Val != val {
			tail.Next = &ListNode{Val: head.Val, Next: nil}
			tail = tail.Next
		}
		head = head.Next
	}
	return root.Next

}

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	removeElements(l1, 4)
}
