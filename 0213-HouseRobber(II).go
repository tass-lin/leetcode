package main

import "fmt"

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	return max(robbing(nums[1:]),robbing(nums[:size-1]))
}

func robbing(nums []int) int {
	var odd, even int
	for index, element := range nums {
		if index % 2 == 0 {
			even = max(even+element,odd)
		} else {
			odd = max(even,odd+element)
		}
	}
	return max(even,odd)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(rob([]int{2,3,2})) //ans:3
	//fmt.Println(rob([]int{1,2,3,1})) //ans:4
	//fmt.Println(rob([]int{1,2,3,1,3})) //ans:6
	//fmt.Println(rob([]int{1})) //ans:1
	fmt.Println(rob([]int{2,1,1,2})) //ans:3
}

//Example 1:
//
//Input: [2,3,2]
//Output: 3
//Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
//because they are adjacent houses.

//Example 2:
//
//Input: [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//Total amount you can rob = 1 + 3 = 4.