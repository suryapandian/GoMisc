package main

import (
	"fmt"
	"math"
)

var maxElement float64

func init() {
	maxElement = math.Pow(10, 9)
}

func main() {
	var n int
	fmt.Println("Enter the number of elements in the simulator")
	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the element in position", i)
		fmt.Scanf("%v", &a[i])
	}
	fmt.Println("Minimum difference", minDifference(a))
}

func minDifference(a []int) (r int) {
	r = int(maxElement)
	n := len(a)
	for i, e := range a {
		for j := i + 1; j < n; j++ {
			d := (e - a[j])
			if d < (a[j] - e) {
				d = (a[j] - e)
			}

			if r > d {
				r = d
			}
		}
	}
	return
}
