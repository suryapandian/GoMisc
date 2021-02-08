package main

func main() {
    fmt.Println("Nicely done Surya")
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // result:= addTwoNumbersWithCarry(l1,l2,0)
    // return result
    r := &ListNode{}
    result := r

    for true {

        if l1 == nil && l2 == nil {
            break
        }
        if l1 != nil {
            result.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            result.Val += l2.Val
            l2 = l2.Next
        }

        carry := result.Val / 10
        result.Val = result.Val % 10

        if l1 != nil || l2 != nil || carry != 0 {
            result.Next = &ListNode{}
            result = result.Next
            result.Val = carry
        }
    }
    return r
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
    result := &ListNode{}
    if l1 == nil && l2 == nil {
        result.Val += carry
        return result
    }
    if l1 != nil {
        result.Val += l1.Val
        l1 = l1.Next
    }
    if l2 != nil {
        result.Val += l2.Val
        l2 = l2.Next
    }

    result.Val += carry
    carry = result.Val / 10
    result.Val = result.Val % 10

    if l1 != nil || l2 != nil || carry != 0 {
        result.Next = addTwoNumbersWithCarry(l1, l2, carry)
    }

    return result
}
