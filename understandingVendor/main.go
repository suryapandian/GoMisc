package main

import "fmt"

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, v := range nums {
        if j, ok := seen[target-v]; ok {
            return []int{i, j}
        }
        seen[v] = i
    }
    return nil
}

func main() {
    a := []int{3, 3}
    b := twoSum(a, 6)
    fmt.Println(b)
}
