package main

import "fmt"

func removeElement(nums []int, val int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }
    return len(nums)
}

func main() {
    arr := []int{9}
    fmt.Println("result", removeElement(arr))
}
