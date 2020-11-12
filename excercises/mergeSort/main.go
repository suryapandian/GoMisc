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

	result := mergeSort(a)
	fmt.Println("Result array", result)

}

func merge(left, right []int) (result []int) {
	i, j := 0, 0
	for (i < len(left)) && (j < len(right)) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result

}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a

	}
	middle := len(a) / 2
	left := mergeSort(a[:middle])
	right := mergeSort(a[middle:])
	return merge(left, right)
}
