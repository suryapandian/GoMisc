package main

import "fmt"

func quickSort(arr []int) {
	n := len(arr)
	if n == 0 || n == 1 {
		return
	}

	if n == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return
	}
	if n == 3 {

	}
	return

}

func main() {
	arr := []int{4, 2}
	quickSort(arr)
	fmt.Println("Sorted aarrry", arr)
}
