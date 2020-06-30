package main

import (
	"sort"
)

func singleNumber(nums []int) int {
	size := len(nums)
	sort.Ints(nums)

	for i := 0; i < size -1; i+=3 {
		if i >= size {
			break
		}
		temp := nums[i:i+2]
		if temp[0] != temp[1] {
			return temp[0]
		}
	}
	return nums[size-1]
}

func main()  {
	println(singleNumber([]int{2,2,3,2})) //3
	println(singleNumber([]int{0,1,0,1,0,1,99})) //99
	println(singleNumber([]int{0,1,0,1,0,1,5,5,5,3})) //3
}

//https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0137.single-number-ii/single-number-ii.go 強

//Example 1:
//
//Input: [2,2,3,2]
//Output: 3

//Example 2:
//
//Input: [0,1,0,1,0,1,99]
//Output: 99