package main    

import "fmt"

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, v := range nums {
        requiredNumber := target - v
        if requiredNumber < 0 {
            requiredNumber *= -1
        }

        seen[v] = i
        if j, ok := seen[requiredNumber]; ok  {
            return []int{i, j}
        }
    }
    return nil
}

func main(){
    a := []int{3,3}
    b := twoSum(a,6)
   fmt.Println(b)
}
