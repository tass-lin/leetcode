package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	temp1 := sortOut(nums1)
	temp2 := sortOut(nums2)

	if len(temp1) > len(temp2) {
		temp1, temp2 = temp2, temp1
	}

	for index := range temp1 {
		if temp2[index] {
			res = append(res,index)
		}
	}
	return res
}

func sortOut(arr []int) map[int]bool  {
	m := make(map[int]bool,len(arr))

	for _, element := range arr {
		m[element] = true
	}
	return m
}

func main() {
	fmt.Printf("%v\n",intersection([]int{1,2,2,1},[]int{2,2})) // [2]
	fmt.Printf("%v\n",intersection([]int{4,9,51},[]int{9,4,9,8,4})) // [9,4]
}

// Given two arrays, write a function to compute their intersection.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]

// Note:
// Each element in the result must be unique.
// The result can be in any order.