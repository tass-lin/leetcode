package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i] - 1]{
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := []int{}
	for index, element := range nums {
		if element != index + 1 {
			res = append(res,index+1)
		}
	}
	return res
}

func findDisappearedNumbers1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[nums[i] - 1]{
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	res := []int{}
	for index, element := range nums {
		if element != index + 1 {
			res = append(res,index+1)
		}
	}
	return res
}

func main() {
	fmt.Printf("%v",findDisappearedNumbers([]int{4,3,2,7,8,2,3,1})) // [5,6]
	fmt.Printf("%v",findDisappearedNumbers1([]int{4,3,2,7,8,2,3,1})) // [5,6]
}

// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements of [1, n] inclusive that do not appear in this array.
// Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//
// Example:
// Input:
// [4,3,2,7,8,2,3,1]
// Output:
// [5,6]