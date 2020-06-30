package main

import "fmt"

var nums []int
func rotate(nums []int, k int)  {
	size := len(nums)
	if k > size {
		k %= size
	}
	if k == 0 || k == size {
		return
	}
	for i := 0; i < k; i++ {
		for j := 1; j < size; j++ {
			nums[0], nums[j] = nums[j], nums[0]
		}
	}
}

func rotate1(nums []int, k int)  {
	size := len(nums)
	if k > size {
		k %= size
	}
	if k == 0 || k == size {
		return
	}

	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, i, j int)  {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate2(nums []int, k int)  {
	size := len(nums)
	if k > size {
		k %= size
	}
	if k == 0 || k == size {
		return
	}

	temp := nums

	nums = append(nums[size-k:], nums[:size-k]...)
	for i := 0; i < len(temp); i++ {
		temp[i] = nums[i]
	}
}

func main() {
	nums = []int{1,2,3,4,5,6,7}
	rotate(nums,3)
	fmt.Printf("%v",nums)  // [5,6,7,1,2,3,4]

	nums = []int{-1,-100,3,99}
	rotate(nums,2)
	fmt.Printf("%v",nums) // [3,99,-1,-100]


	nums = []int{1,2,3,4,5,6,7}
	rotate1(nums,3)
	fmt.Printf("%v",nums)  // [5,6,7,1,2,3,4]

	nums = []int{-1,-100,3,99}
	rotate1(nums,2)
	fmt.Printf("%v",nums) // [3,99,-1,-100]


	nums = []int{1,2,3,4,5,6,7}
	rotate2(nums,3)
	fmt.Printf("%v",nums)  // [5,6,7,1,2,3,4]

	nums = []int{-1,-100,3,99}
	rotate2(nums,2)
	fmt.Printf("%v",nums) // [3,99,-1,-100]
}

//Given an array, rotate the array to the right by k steps, where k is non-negative.
//
//Follow up:
//
//Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
//Could you do it in-place with O(1) extra space?

//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]

//Input: nums = [-1,-100,3,99], k = 2
//Output: [3,99,-1,-100]
//Explanation:
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]