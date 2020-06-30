package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	max, temp := 1, 1
	if len(nums) == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1]{
			temp++
		} else {
			max = maxFunc(max, temp)
			temp = 1
		}
	}
	return maxFunc(max, temp)
}
func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7})) // 3
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2,2})) // 1
	fmt.Println(findLengthOfLCIS([]int{1,3,5,7})) // 4
}

// Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).
//
// Example 1:
// Input: [1,3,5,4,7]
// Output: 3
// Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
// Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
//
// Example 2:
// Input: [2,2,2,2,2]
// Output: 1
//
// Explanation: The longest continuous increasing subsequence is [2], its length is 1.
// Note: Length of the array will not exceed 10,000.