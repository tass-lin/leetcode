package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, element := range nums {
		if element == max1 || element == max2 {
			continue
		}
		switch {
		case max1 < element:
			max3, max2, max1 = max2, max1, element
		case max2 < element:
			max3, max2 = max2, element
		case max3 < element:
			max3 = element
		}
	}

	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1})) // 1
	fmt.Println(thirdMax([]int{1, 2})) // 2
	fmt.Println(thirdMax([]int{2, 2, 3, 1})) // 1
}

// Given a non-empty array of integers, return the third maximum number in this array. If it does not exist,
// return the maximum number. The time complexity must be in O(n).
//
// Example 1:
// Input: [3, 2, 1]
// Output: 1
// Explanation: The third maximum is 1.
//
// Example 2:
// Input: [1, 2]
// Output: 2
// Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//
// Example 3:
// Input: [2, 2, 3, 1]
// Output: 1
// Explanation: Note that the third maximum here means the third maximum distinct number.
// Both numbers with value 2 are both considered as second maximum.