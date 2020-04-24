package main    

import "fmt"

func twoSum(nums []int, target int) []int {
    for i:=0; i<len(nums)-1; i++ {
      for j:=i+1; j<len(nums); j++ {
        if nums[i] + nums[j] == target {
          return []int{i, j}
        }
      }
    }
    return nil
}
func main(){
    a := []int{3,3}
    b := twoSum(a,6)
   fmt.Println(b)


}