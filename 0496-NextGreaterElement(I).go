package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int,len(nums1))
	mNum2 := make(map[int]int,len(nums2))
	for index, element := range nums2 {
		mNum2[element] = index
	}

	for index, element := range nums1 {
		res[index] = -1
		for j := mNum2[element] + 1; j < len(nums2); j++ {
			if element < nums2[j] {
				res[index] = nums2[j]
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v",nextGreaterElement([]int{4,1,2},[]int{1,3,4,2})) // [-1,3,-1]
	fmt.Printf("%v",nextGreaterElement([]int{2,4},[]int{1,2,3,4})) // [3,-1]
	fmt.Printf("%v",nextGreaterElement([]int{4,1,2},[]int{1,2,3,4})) // [-1,2,3]
}


// You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2.
// Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
// The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2.
// If it does not exist, output -1 for this number.
//
// Example 1:
// Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
// Output: [-1,3,-1]
// Explanation:
// For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
// For number 1 in the first array, the next greater number for it in the second array is 3.
// For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
//
// Example 2:
// Input: nums1 = [2,4], nums2 = [1,2,3,4].
// Output: [3,-1]
// Explanation:
// For number 2 in the first array, the next greater number for it in the second array is 3.
// For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
//
// Note:
// All elements in nums1 and nums2 are unique.
// The length of both nums1 and nums2 would not exceed 1000.