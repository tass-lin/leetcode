package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	size := len(nums)
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < size; i++ {
		if nums[i-1]+1 != nums[i] {
			return nums[i-1]+1
		}
	}
	return nums[size-1]+1
}

//^：按位異或運算符，是雙目運算符。其功能是參與運算的兩數各對應的二進位相異或，當兩對應的二進位相異時，結果為1。
//特別地，^支持交換律
//1^1^2^2^3^3 == 1^2^2^3^3^1 == 1^3^2^3^1^2
func missingNumber1(nums []int) int {
	xor := 0
	for i, n := range nums {
		xor ^= i ^ n
	}

	//所有的i再加上len（nums），就相當於0,1，...，n。
	//利用相同的數異或為0，及其交換律
	//xor最後的值，就是那個缺失的數

	return xor ^ len(nums)
}

func main() {
	fmt.Println(missingNumber([]int{3,0,1})) //2
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1})) //8
	fmt.Println(missingNumber([]int{1,2})) //0

	fmt.Println(missingNumber1([]int{3,0,1})) //2
	fmt.Println(missingNumber1([]int{9,6,4,2,3,5,7,0,1})) //8
	fmt.Println(missingNumber1([]int{1,2})) //0
}

//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
//Example 1:
//
//Input: [3,0,1]
//Output: 2
//Example 2:
//
//Input: [9,6,4,2,3,5,7,0,1]
//Output: 8
//Note:
//Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?