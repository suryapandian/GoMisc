package main

import (
	"fmt"
	"math"
)

func rotate(matrix [][]int) {
	n := len(matrix[0])
	compN := n - 1
	for i := 0; i <= n/2; i++ {
		fmt.Println(i)
		for j := i; j < (compN - i); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[compN-j][i]
			matrix[compN-j][i] = matrix[compN-i][compN-j]
			matrix[compN-i][compN-j] = matrix[j][compN-i]
			matrix[j][compN-i] = temp
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}, {13, 14, 15, 16}})
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {13, 14, 15}})
}
