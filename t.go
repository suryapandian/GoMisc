/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) (result *ListNode) {
    for _, list := range lists {
        result = mergeLists(result, list)
    }
    return result

}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    r := &ListNode{}
    result := r
    for true {
        if l1 == nil && l2 == nil {
            break
        }

        if l1 != nil {
            if l2 == nil {
                result.Val = l1.Val
                l1 = l1.Next
            } else {
                if l1.Val >= l2.Val {
                    result.Val = l2.Val
                    l2 = l2.Next
                } else {
                    result.Val = l1.Val
                    l1 = l1.Next
                }

            }
        } else {
            if l2 != nil {
                result.Val = l2.Val
                l2 = l2.Next
            }
        }

        if l1 != nil || l2 != nil {
            result.Next = &ListNode{}
            result = result.Next
        }

    }
    return r
}
