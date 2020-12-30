package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) (ans int) {
	n := len(nums)
	var numMap = make(map[int]int, n)
	for _, ele := range nums {
		numMap[ele] = 1
	}
	for i := 1; i <= n; i++ {
		if _, ok := numMap[i]; !ok {
			return i
		}
	}
	return n + 1
}

func main() {
	arr := []int{-1, -1, 5, 7}
	fmt.Printf("First missing positive number for arr %v is : %v \n", arr, firstMissingPositive(arr))
	arr = []int{-1, -1, 0, 1}
	fmt.Printf("First missing positive number for arr %v is : %v \n", arr, firstMissingPositive(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Printf("First missing positive number for arr %v is : %v \n", arr, firstMissingPositive(arr))

	arr = []int{-1, -1, 99, 100}
	fmt.Printf("First missing positive number for arr %v is : %v \n", arr, firstMissingPositive(arr))
}
