package main

import "math"

func maxSubArray(nums []int) int {
	nlen := len(nums)
	max := math.MinInt32

	for index, _ := range nums {
		buffer := 0
		for i := index; i <= nlen - 1 ; i++ {
			buffer += nums[i]
			if max < buffer {
				max = buffer
			}
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
	sum, maxSum := math.MinInt32, math.MinInt32
	for _, n := range nums {
		// sum+n < n，不如從n開始
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	println(maxSubArray([]int{1}))
	println(maxSubArray([]int{-2,-1}))

	println(maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4}))
	println(maxSubArray1([]int{1}))
	println(maxSubArray1([]int{-2,-1}))
}
