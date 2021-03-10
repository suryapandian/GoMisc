package main

import (
    "fmt"
)

func subarraySum(a []int, k int) int {
    res, sum := 0, 0
    rec := make(map[int]int, len(a)) //we are storing sum of continuous array in this map
    rec[0] = 1                       //To account for cases when internal sum == k

    for i := range a {

        sum += a[i]

        res += rec[sum-k]

        rec[sum]++
    }

    return res
}

func main() {
    arr := []int{1, 11, 5, 5}
    k := 10
    result := subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

    arr = []int{1, 1, 1}
    k = 2
    result = subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

    arr = []int{1, 2, 3}
    k = 3
    result = subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

    arr = []int{1}
    k = 1
    result = subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

    arr = []int{1, -1, 0}
    k = 0
    result = subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

    arr = []int{1, 2, 1, 2, 1}
    k = 3
    result = subarraySum(arr, k)
    fmt.Printf("solution for arr: %v, sum: %d = %d \n", arr, k, result)

}
