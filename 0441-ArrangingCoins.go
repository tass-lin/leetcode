package main

import "fmt"

func arrangeCoins(n int) int {
	stair, max := 1, 0
	for n > 0 {
		max = stair
		n -= stair
		if n < 0 {
			max = stair - 1
		}
		stair++
	}
	return max
}

func main() {
	fmt.Println(arrangeCoins(5)) // 2
	fmt.Println(arrangeCoins(8)) // 3
}

// You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.
// Given n, find the total number of full staircase rows that can be formed.
// n is a non-negative integer and fits within the range of a 32-bit signed integer.
//
// Example 1:
// n = 5
// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤
// Because the 3rd row is incomplete, we return 2.
//
// Example 2:
// n = 8
// The coins can form the following rows:
// ¤
// ¤ ¤
// ¤ ¤ ¤
// ¤ ¤
// Because the 4th row is incomplete, we return 3.