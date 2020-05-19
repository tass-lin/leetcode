package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max, temp := 0, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			temp = i
		} else {
			max = mathMax(max,i - temp)
		}
	}
	return max
}

func mathMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes1(nums []int) int {
	max := 0

	for i, temp := 0, -1; i < len(nums); i++ {
		if nums[i] == 0 {
			temp = i
		} else {
			if max < i - temp {
				max = i - temp
			}
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1})) // 3
	fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0,1})) // 2
}

// Given a binary array, find the maximum number of consecutive 1s in this array.
// Example 1:
// Input: [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s.
// The maximum number of consecutive 1s is 3.
//
// Note:
// The input array will only contain 0 and 1.
// The length of input array is a positive integer and will not exceed 10,000