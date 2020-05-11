package main

import (
	"fmt"
)

func rob(nums []int) int {
	var odd, even int
	for index, element := range nums {
		if index % 2 == 0 {
			even = max(even+element, odd)
		} else {
			odd = max(even, odd+element)
		}
		println(even,odd)
	}
	return max(even, odd)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2,1,1,2})) //4
	//fmt.Println(rob([]int{1,2,3,1})) //ans:4
	//fmt.Println(rob([]int{1,5,3,1})) //ans:6
	//fmt.Println(rob([]int{2,7,9,3,1})) //12
	//fmt.Println(rob([]int{1,3,1})) //3
	//fmt.Println(rob([]int{1,2,1,1})) //3

}

//Example 1:
//
//Input: [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//Total amount you can rob = 1 + 3 = 4.

//Example 2:
//
//Input: [2,7,9,3,1]
//Output: 12
//Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
//Total amount you can rob = 2 + 9 + 1 = 12.