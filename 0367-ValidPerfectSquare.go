package main

import "fmt"

func isPerfectSquare(num int) bool {
	r := intSqrt(num)
	return r*r == num
}

func intSqrt(n int) int {
	r := n

	for r*r > n {
		r = (r + n/r) / 2
	}

	return r
}

func main() {
	fmt.Println(isPerfectSquare(16)) // true
	// fmt.Println(isPerfectSquare(14)) // false
}

// Given a positive integer num, write a function which returns True if num is a perfect square else False.
// Note: Do not use any built-in library function such as sqrt.
//
// Example 1:
// Input: 16
// Output: true
//
// Example 2:
// Input: 14
// Output: false