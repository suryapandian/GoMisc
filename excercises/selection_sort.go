package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min, index := arr[i], i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min, index = arr[j], j
			}
		}
		if arr[i] > min {
			arr[index] = arr[i]
			arr[i] = min
		}
	}
}

func main() {
	input := []int{9, 34, 56, 10, 83, 6, 55, 10, 0, -1, 34, 57, -86}
	selectionSort(input)
	fmt.Println("Sorting by selection sort", input)
}
