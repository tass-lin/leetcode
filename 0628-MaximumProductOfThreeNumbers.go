package main

import (
	"fmt"
)

func maximumProduct(nums []int) int {
	max := -1001
	max1 := -1001
	max2 := -1001
	min1 := 1001
	min2 := 1001

	for _, element := range nums {
		switch {
		case element > max:
			max2, max1, max = max1, max, element
		case element > max1:
			max2, max1 = max1, element
		case element > max2:
			max2 = element
		}

		switch {
		case element < min1:
			min2, min1 = min1, element
		case element < min2:
			min2 = element
		}
	}

	return maxFunc(max1 * max2, min1 * min2) * max
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumProduct([]int{1,2,3})) // 6
	fmt.Println(maximumProduct([]int{1,2,3,4})) // 24
	fmt.Println(maximumProduct([]int{-4,-3,-2,-1,60})) // 720
}

// Given an integer array, find three numbers whose product is maximum and output the maximum product.
//
// Example 1:
// Input: [1,2,3]
// Output: 6
//
//
// Example 2:
// Input: [1,2,3,4]
// Output: 24
//
// Note:
// The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
// Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.