package main

import (
    "fmt"
    "sort"
)

func subarraySum(nums []int, k int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }

    if n == 1 {
        if nums[0] == k {
            return 1
        }
        return 0
    }

    var result int

    for i := 1; i < n; i++ {
        for j := 0; j+i <= n; j++ {
            if isSumEqual(nums[j:j+i], k) {
                result++
            }
        }
    }

    if isSumEqual(nums, k) {
        result++
    }

    return result
}

func isSumEqual(arr []int, sum int) bool {
    nums := make([]int, len(arr))
    copy(nums, arr)
    sort.Ints(nums)
    arraySum := 0
    for _, e := range nums {
        arraySum += e
        if arraySum > sum {
            return false
        }
    }
    return arraySum == sum
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
