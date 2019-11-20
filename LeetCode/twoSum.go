/* Problem Statement

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

func twoSum(nums []int, target int) []int {
    var results []int
    m := make(map[int]int)

    for i, v := range nums {
        expectedNum := getOtherNumber(v,target)
        m[v] = i+1
        if m[expectedNum]!= 0{
          return []int{i, m[expectedNum]-1}  
        }
       
    }
    return results
}

func getOtherNumber(num, target int) (result int){
    result = target - num
    if result < 0 {
            result = result * (-1)
    }
    return
}
