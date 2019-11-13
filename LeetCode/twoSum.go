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
    for i, v := range nums {
        expectedNum := target - v
        if expectedNum < 0 {
            expectedNum = expectedNum * (-1)
        }
        indexOfExpectedNum := contains(nums, expectedNum)
        if nums[i] + nums[indexOfExpectedNum] == target{
            return []int{i, indexOfExpectedNum}
        }
    }
    return results
}


func contains(array []int, num int) int{
   for i,v := range array {
        if v == num {
            return i
        }
  } 
  return 0
}