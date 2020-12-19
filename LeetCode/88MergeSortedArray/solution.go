package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}

	for i1, i, j := 0, 0, 0; i < m+n && j < n; i++ {
		if i1 >= m || nums2[j] < nums1[i] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
		} else {
			i1++
		}
	}

	return
}

func main() {
	arr1 := make([]int, 8)
	arr1[0], arr1[1], arr1[2] = 2, 4, 9
	//arr1 := []int{2, 4, 9}
	arr2 := []int{1, 3, 6, 7, 10}
	merge(arr1, 3, arr2, 5)
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("rsee", arr1)

}
