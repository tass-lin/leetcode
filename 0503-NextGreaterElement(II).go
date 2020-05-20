package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	res := make([]int,len(nums))

	for index, element := range nums {
		res[index] = -1
		i, j := index + 1 , 0
		for i < len(nums){
			if nums[i] > element{
				res[index] = nums[i]
				break
			}
			i++
		}
		if i == len(nums) {
			for j <= index {
				if nums[j] > element  {
					res[index] = nums[j]
					break
				}
				j++
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n",nextGreaterElements([]int{-1,0,-2})) // [0,-1,-1]
	fmt.Printf("%v\n",nextGreaterElements([]int{1,2,1})) // [2,-1,2]
	fmt.Printf("%v\n",nextGreaterElements([]int{1,2,3,4,5})) // [2,3,4,5,-1]
	fmt.Printf("%v\n",nextGreaterElements([]int{5,4,3,2,1})) // [-1,5,5,5,5]
	fmt.Printf("%v\n",nextGreaterElements([]int{1,5,3,6,8})) // [5,6,6,8,-1]
	fmt.Printf("%v\n",nextGreaterElements([]int{100,1,11,1,120,111,123,1,-1,-100})) // [120,11,120,120,123,123,-1,100,100,100]
	fmt.Printf("%v\n",nextGreaterElements([]int{1,8,-1,-100,-1,222,1111111,-111111})) // [8,222,222,-1,222,1111111,-1,1]
}


// Given a circular array (the next element of the last element is the first element of the array),
// print the Next Greater Number for every element.
// The Next Greater Number of a number x is the first greater number to its traversing-order next in the array,
// which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.
//
// Example 1:
// Input: [1,2,1]
// Output: [2,-1,2]
// Explanation: The first 1's next greater number is 2;
// The number 2 can't find next greater number;
// The second 1's next greater number needs to search circularly, which is also 2.
// Note: The length of given array won't exceed 10000.