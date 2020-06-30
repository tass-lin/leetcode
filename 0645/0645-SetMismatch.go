package main

import (
	"fmt"
	"math"
)

func findErrorNums(nums []int) []int {
	m := make(map[int]bool,len(nums))
	var res []int
	for _, element := range nums {
		if m[element] == true {
			res = append(res, element)
		}
		m[element] = true
	}
	temp := math.MaxInt32
	for _, element := range nums {
		if m[1] != true {
			temp = 1
			break
		}
		if m[element+1] != true {
			temp = min(temp, element+1)
		}
	}
	return append(res, temp)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func findErrorNums1(nums []int) []int {
	arr := make([]int,len(nums))
	res := make([]int,2)

	for _, element := range nums {
		arr[element-1]++
	}

	for index, element := range arr {
		if element == 2 {
			res[0] = index + 1
		} else if element == 0 {
			res[1] = index + 1
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n", findErrorNums([]int{1,2,2,4})) // [2,3]
	fmt.Printf("%v\n", findErrorNums([]int{1,1})) // [1,2]
	fmt.Printf("%v\n", findErrorNums([]int{2,2})) // [2,1]
	fmt.Printf("%v\n", findErrorNums([]int{3,2,2})) // [2,1]
	fmt.Printf("%v\n", findErrorNums([]int{3,3,1})) // [3,2]

	fmt.Printf("%v\n", findErrorNums1([]int{1,2,2,4})) // [2,3]
	fmt.Printf("%v\n", findErrorNums1([]int{1,1})) // [1,2]
	fmt.Printf("%v\n", findErrorNums1([]int{2,2})) // [2,1]
	fmt.Printf("%v\n", findErrorNums1([]int{3,2,2})) // [2,1]
	fmt.Printf("%v\n", findErrorNums1([]int{3,3,1})) // [3,2]
	fmt.Printf("%v\n", findErrorNums1([]int{2,3,2})) // [2,1]
	
}

// The set S originally contains numbers from 1 to n.
// But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set,
// which results in repetition of one number and loss of another number.
//
// Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.
//
// Example 1:
// Input: nums = [1,2,2,4]
// Output: [2,3]
//
// Note:
// The given array size will in the range [2, 10000].
// The given array's numbers won't have any order.