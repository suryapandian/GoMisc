package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix[0])
	compN := n - 1
	for outer := 0; outer <= n/2; outer++ {
		fmt.Println(outer)
		for inner := outer; inner < (compN - outer); inner++ {
			temp := matrix[outer][inner]
			matrix[outer][inner] = matrix[compN-inner][outer]
			matrix[compN-inner][outer] = matrix[compN-outer][compN-inner]
			matrix[compN-outer][compN-inner] = matrix[inner][compN-outer]
			matrix[inner][compN-outer] = temp
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}, {13, 14, 15, 16}})
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {13, 14, 15}})
}
