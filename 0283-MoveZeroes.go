package main

import "fmt"

var nums []int
func moveZeroes(nums []int)  {
	size, i, j := len(nums), 0, 0

	for j < size {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < size {
		nums[i] = 0
		i++
	}
}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Printf("%v",nums) // [1,3,12,0,0]
}


//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Example:
//
//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
//Note:
//
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.