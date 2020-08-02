package main

import "fmt"

func main() {
	// var n int
	// fmt.Println("Enter the number of elements in the array")
	// fmt.Scanf("%d", &n)

	// a := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Println("Enter the element in position", i)
	// 	fmt.Scanf("%v", &a[i])
	// }

	a := [5]int{45, 5, 334, 9, 100}
	n := 5
	var result int
	for i := 0; i < n; i++ {
		ra := rotate(a, i)
		numSwap := numOfSwap(ra)
		fmt.Println("Number of swap", numSwap, ra)
		if (i == 0) || (result > numSwap) {
			result = numSwap
		}
	}

	fmt.Println("Minimum number of swap", result)

}

func numOfSwap(a [5]int) (n int) {
	var tmp, j int
	for i := 1; i < len(a); i++ {
		tmp = a[i]
		for j = i; (j > 0) && (a[j-1] > a[j]); j-- {
			a[j] = a[j-1]
			a[j-1] = tmp
			n++
		}
	}
	return n
}

func rotate(a [5]int, n int) [5]int {
	for n >= 0 {
		var tmp, tmp2 int
		tmp = a[0]
		for i := 1; i < len(a); i++ {
			tmp2 = a[i]
			a[i] = tmp
			tmp = tmp2
		}
		a[0] = tmp
		n--
	}
	return a

}
