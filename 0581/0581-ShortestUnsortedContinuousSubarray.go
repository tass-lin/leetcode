package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	left, right := 0, -1
	min, max := math.MaxInt32, math.MinInt32

	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] >= max {
			max = nums[i]
		} else {
			right = i
		}

		if nums[j] <= min {
			min = nums[j]
		} else {
			left = j
		}
	}
	fmt.Println(right,left)

	return right - left + 1
}

func main() {
	fmt.Printf("%v\n", findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})) // 5
	fmt.Printf("%v\n", findUnsortedSubarray([]int{1,3,2,2,2})) // 4
	fmt.Printf("%v\n", findUnsortedSubarray([]int{1,2,3,3,3})) // 0
}

// Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order,
// then the whole array will be sorted in ascending order, too.
//
// You need to find the shortest such subarray and output its length.

// Example 1:
// Input: [2, 6, 4, 8, 10, 9, 15]
// Output: 5
// Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
// Note:
// Then length of the input array is in range [1, 10,000].
// The input array may contain duplicates, so ascending order here means <=.