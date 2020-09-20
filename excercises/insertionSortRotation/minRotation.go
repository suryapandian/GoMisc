package main

import "fmt"

func main() {
	a := [5]int{45, 5, 334, 9, 100}
	n := len(a)
	var result int
	for i := 0; i < n; i++ {
		ra := rotate(a, i)
		fmt.Println("after rotate", i, a)
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
	if n > 0 {
		b := a
		for i := 0; i < len(a); i++ {
			moveBy := (i + (n % len(a))) % len(a)
			a[moveBy] = b[i]
		}
	}
	return a

}
