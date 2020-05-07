package main

import "sort"

func singleNumber(nums []int) int {
	size := len(nums)
	sort.Ints(nums)
	if size == 1{
		return nums[0]
	}
	for i := 0; i < size - 2; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		i++
	}

	return nums[size-1]
}

func main()  {
	println(singleNumber([]int{2,2,1})) //1
	println(singleNumber([]int{4,1,2,1,2})) //4
}
// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0136.single-number/single-number.go 有空再研究

//Example 1:
//
//Input: [2,2,1]
//Output: 1

//Example 2:
//
//Input: [4,1,2,1,2]
//Output: 4