package main

import "fmt"

func main() {
	var n, m int
	fmt.Println("Enter the value of m")
	fmt.Scanf("%d", &m)
	fmt.Println("Enter the number of elements in the array")
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the element in position", i)
		fmt.Scanf("%v", &a[i])
	}

	fmt.Println("Number of iterations", bubbleSort(a, m))

}

func bubbleSort(a []int, numIterations int) (r int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < (len(a) - 1); j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				fmt.Println(i, numIterations)
				if i >= numIterations {
					r++
				}
			}
		}
	}
	return
}
