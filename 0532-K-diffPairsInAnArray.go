package main

import "fmt"

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m, res := make(map[int]int,len(nums)), 0
	
	for _, element := range nums {
		m[element]++
	}
	if k == 0 {
		for _, element := range m {
			if element > 1 {
				res++
			}
		}
		return res
	}
	for index, _ := range m {
		if m[index - k] > 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2)) // 2
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1)) // 4
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0)) // 1
	fmt.Println(findPairs([]int{1, 1, 1, 1, 1}, 0)) // 1
}

// Given an array of integers and an integer k, you need to find the number of unique k-diff pairs in the array.
// Here a k-diff pair is defined as an integer pair (i, j), where i and j are both numbers in the array and their absolute difference is k.
//
// Example 1:
// Input: [3, 1, 4, 1, 5], k = 2
// Output: 2
// Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
// Although we have two 1s in the input, we should only return the number of unique pairs.
//
// Example 2:
// Input:[1, 2, 3, 4, 5], k = 1
// Output: 4
// Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
//
// Example 3:
// Input: [1, 3, 1, 5, 4], k = 0
// Output: 1
// Explanation: There is one 0-diff pair in the array, (1, 1).
//
// Note:
// The pairs (i, j) and (j, i) count as the same pair.
// The length of the array won't exceed 10,000.
// All the integers in the given input belong to the range: [-1e7, 1e7].