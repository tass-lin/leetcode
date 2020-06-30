package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int,len(nums))
	res := 0
	for _, element := range nums {
		m[element] += 1
	}

	for index, element := range m {
		element2, ok := m[index+1]
		if ok {
			temp := element + element2
			if res < temp {
				res = temp
			}
		}
	}

	return res
}

func main() {
	fmt.Printf("%v", findLHS([]int{1,3,2,2,5,2,3,7})) // 5 [3,2,2,2,3].
	fmt.Printf("%v", findLHS([]int{1,2,3,4})) // 2
}

// We define a harmounious array as an array where the difference between its maximum value and its minimum value is exactly 1.
// Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.
//
// Example 1:
// Input: [1,3,2,2,5,2,3,7]
// Output: 5
// Explanation: The longest harmonious subsequence is [3,2,2,2,3].
//
// Note: The length of the input array will not exceed 20,000.