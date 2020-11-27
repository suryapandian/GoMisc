package main

import "fmt"

func sumOfArray(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	return arr[0] + sumOfArray(arr[1:])
}

func max(arr []int) int {
	n := len(arr)
	if n == 2 {
		if arr[0] < arr[1] {
			return arr[1]
		} else {
			return arr[0]
		}
	}

	restMax := max(arr[1:])

	if arr[0] > restMax {
		return arr[0]
	}
	return restMax
}

func main() {
	input := []int{107, 34, 56, 10, 83, 6, 55, 10, 0, -1, 34, 57, -86}
	fmt.Println("Max of array", max(input))
	fmt.Println("Sum of array", sumOfArray(input))
}
