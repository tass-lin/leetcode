package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	temp1, temp2 := sortOut(nums1), sortOut(nums2)
	if len(temp1) > len(temp2) {
		temp1, temp2 = temp2, temp1
	}
	for index := range temp1 {
		temp1[index] = min(temp1[index] , temp2[index])
	}
	for index, element := range temp1 {
		for i := 0; i < element; i++{
			res = append(res,index)
		}
	}
	return res
}

func sortOut(arr []int) map[int]int {
	m := make(map[int]int, len(arr))
	for _, element := range arr {
		m[element] += 1
	}
	return m
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func main() {
	// fmt.Printf("%v\n",intersect([]int{1,2,2,1},[]int{2,2})) // [2,2]
	fmt.Printf("%v\n",intersect([]int{4,9,5},[]int{9,4,9,8,4})) // [4,9]
}

// Given two arrays, write a function to compute their intersection.
//
// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
//
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
//
// Note:
// Each element in the result should appear as many times as it shows in both arrays.
// The result can be in any order.
//
// Follow up:
// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?