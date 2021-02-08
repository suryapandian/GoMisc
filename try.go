package main

import (
    "fmt"
)

func main() {
    str := []int{5, 2, 4, 6, 3, 7}
    fmt.Printf("Solution to %d is %d \n", str, Solution(str))
    str = []int{1, 23, 12, 1, 56, 7, 8, 9, 1}
    fmt.Printf("Solution to %d is %d \n", str, Solution(str))
    str = []int{5, 73, 12, 7, 56, 7, 8, 9, 1}
    fmt.Printf("Solution to %d is %d \n", str, Solution(str))
    str = []int{5, 3, 12, 2, 56, 7, 8, 9, 1}
    fmt.Printf("Solution to %d is %d \n", str, Solution(str))
    str = []int{5, 53, 12, 11, 56, 7, 8, 9, 1}
    fmt.Printf("Solution to %d is %d \n", str, Solution(str))
}

func Solution(A []int) int {
    n := len(A)
    min := A[1] + A[3]
    for i := 1; i < n-3; i += 2 {
        for j := i + 2; j < n-1; j++ {
            sum := A[i] + A[j]
            //fmt.Println("i j A[i] A[j] sum", i, j, A[i], A[j], sum)
            if sum < min {
                min = sum
            }
        }

    }
    return min
}
