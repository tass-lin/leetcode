package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n % 3 == 0 {
		n /= 3
	}

	if n == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(27)) //true
	fmt.Println(isPowerOfThree(0)) //false
	fmt.Println(isPowerOfThree(9)) //true
	fmt.Println(isPowerOfThree(45)) //false
}

// Given an integer, write a function to determine if it is a power of three.
//
// Example 1:
// Input: 27
// Output: true

// Example 2:
// Input: 0
// Output: false

// Example 3:
// Input: 9
// Output: true

// Example 4:
// Input: 45
// Output: false

// Follow up:
// Could you do it without using any loop / recursion?