package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	temp := 0

	for i := 0; i < k; i++ {
		temp += nums[i]
	}

	max := temp
	for i := k; i < len(nums); i++ {
		temp = temp - nums[i-k] + nums[i]
		max = maxFunc(max, temp)
	}
	return float64(max) / float64(k)
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4)) // 12.75
	fmt.Println(findMaxAverage([]int{0,4,0,3,2}, 1)) // 4
	fmt.Println(findMaxAverage([]int{4,2,1,3,3}, 2)) // 3
}

// Given an array consisting of n integers,
// find the contiguous subarray of given length k that has the maximum average value.
// And you need to output the maximum average value.
//
// Example 1:
// Input: [1,12,-5,-6,50,3], k = 4
// Output: 12.75
// Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
//
// Note:
// 1 <= k <= n <= 30,000.
// Elements of the given array will be in the range [-10,000, 10,000].