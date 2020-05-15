package main

import "fmt"

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}

	for num % 4 == 0 {
		num /= 4
	}

	return num == 1
}

func main() {
	fmt.Println(isPowerOfFour(16)) // true
	fmt.Println(isPowerOfFour(5)) // false
}

// Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
//
// Example 1:
// Input: 16
// Output: true
//
// Example 2:
// Input: 5
// Output: false
// Follow up: Could you solve it without loops/recursion?