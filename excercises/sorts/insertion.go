package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of elements in the array")
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the element in position", i)
		fmt.Scanf("%v", &a[i])
	}

	result := insertionSort(a)
	fmt.Println("Result array", result)

}

func insertionSort(a []int) []int {
	var tmp, j int
	for i := 1; i < len(a); i++ {
		tmp = a[i]
		for j = i; (j > 0) && (a[j-1] > a[j]); j-- {
			a[j] = a[j-1]
			a[j-1] = tmp
		}
	}
	return a
}
