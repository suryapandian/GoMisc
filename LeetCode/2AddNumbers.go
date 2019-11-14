
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNum := getNumber(l1) + getNumber(l2)
	return arrayToList(numToArray(resultNum))
}

func getNumber(l *ListNode) int {
	return arrayToNum(listToArray(l))
}

func listToArray(l *ListNode) (array []int) {
	currentNode := l
	for currentNode != nil {
		array = append(array, currentNode.Val)
		currentNode = currentNode.Next
	}
	return
}

func arrayToList(array []int) (list *ListNode) {
	var l ListNode
	for _, v := range array {
		list = &ListNode{v, &l}
		l = *list
	}
	return
}

func arrayToNum(numberArray []int) (result int) {
	n := len(numberArray) - 1
	for j := range numberArray {
		i := int(j)
		units := math.Pow(10, float64(i))
		result = result + (int(units) * numberArray[n-i])
	}
	return
}

func numToArray(number int) (array []int) {
	currentNumber := number
	for currentNumber > 0 {
		array = append(array, currentNumber%10)
		currentNumber = currentNumber / 10
	}
	return array
}

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(&l1, &l2)
}
