package main

import "fmt"

func checkPossibility(nums []int) bool {
	var temp, temp2 int
	arr := make([]int,len(nums))
	copy(arr, nums)
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1]{
			nums[i+1] = nums[i]
			i = -1
			temp++
		}
	}
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i+1]{
			arr[i] = arr[i+1]
			i = -1
			temp2++
		}
	}

	return temp <= 1 || temp2 <= 1
}


func checkPossibility1(nums []int) bool {
	var temp, min int
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			temp++
			if min > nums[i+1] {
				nums[i+1] = nums[i]
			}
		} else {
			min = nums[i]
		}
	}
	return temp <= 1
}
func main() {
	fmt.Println(checkPossibility([]int{4,2,3})) // true
	fmt.Println(checkPossibility([]int{4,2,1})) // false
	fmt.Println(checkPossibility([]int{3,4,2,3})) // false
	fmt.Println(checkPossibility([]int{2,3,3,2,4})) // true

	fmt.Println(checkPossibility1([]int{4,2,3})) // true
	fmt.Println(checkPossibility1([]int{4,2,1})) // false
	fmt.Println(checkPossibility1([]int{3,4,2,3})) // false
	fmt.Println(checkPossibility1([]int{2,3,3,2,4})) // true
}

// Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.
// We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
//
// Example 1:
// Input: nums = [4,2,3]
// Output: true
// Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
//
// Example 2:
// Input: nums = [4,2,1]
// Output: false
// Explanation: You can't get a non-decreasing array by modify at most one element.