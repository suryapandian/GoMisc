package main

import "fmt"

func binarySearch(arr []int, num int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	if n == 1 {
		if num == arr[0] {
			return true
		} else {
			return false
		}
	}

	if arr[n/2] == num {
		return true
	}
	if arr[n/2] > num {
		return binarySearch(arr[:n/2], num)
	}

	if arr[n/2] < num {
		return binarySearch(arr[n/2:], num)
	}
	return false
}

func binarySearchIndex(arr []int, num int) (int, bool) {
	low, high := 0, len(arr)
	for low < high {
		a := arr[low:high]
		n := len(a)
		//	fmt.Println("low", "high", low, high, a)
		if a[n/2] == num {
			return low + n/2, true
		}

		if a[n/2] < num {
			low = low + high/2
		}

		if a[n/2] > num {
			high = high - high/2
		}
	}
	return 0, false
}

func main() {
	a := []int{1, 2, 5, 7, 8, 20, 22}
	fmt.Println("looking for 4", binarySearch(a, 4))
	fmt.Println("looking for 8", binarySearch(a, 8))
	fmt.Println("looking for 100", binarySearch(a, 100))
	fmt.Println("looking for 20", binarySearch(a, 20))
	index, pres := binarySearchIndex(a, 2)
	fmt.Printf("looking for 2, present %v, index %d \n", pres, index)
	index, pres = binarySearchIndex(a, 1)
	fmt.Printf("looking for 1, present %v, index %d \n", pres, index)
	index, pres = binarySearchIndex(a, 100)
	fmt.Printf("looking for 100, present %v, index %d \n", pres, index)
	index, pres = binarySearchIndex(a, 22)
	fmt.Printf("looking for 22, present %v, index %d \n", pres, index)

}
