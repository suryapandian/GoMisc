/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    temp := &ListNode{}
    result := temp
    seen := make(map[int]bool)

    for {
        if head == nil {
            break
        }
        val := head.Val
        head = head.Next
        if seen[val] {
            continue
        }
        seen[val] = true
        temp.Next = &ListNode{Val: val}
        temp = temp.Next
    }
    return result.Next
}