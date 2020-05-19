package main

import "fmt"

func findComplement(num int) int {
	temp, res := num, 0

	for temp > 0 {
		res = res * 2 + 1
		temp = temp >> 1
	}

	return res ^ num
}

func findComplement1(num int) int {
	temp, res := num, 0

	for temp > 0 {
		temp >>= 1
		res <<= 1
		res++
	}

	return res ^ num
}

func main() {
	// fmt.Println(findComplement(5)) // 2
	// fmt.Println(findComplement(1)) // 0

	fmt.Println(findComplement1(5)) // 2
	// fmt.Println(findComplement1(1)) // 0
}

// Given a positive integer num, output its complement number. The complement strategy is to flip the bits of its binary representation.
//
// Example 1:
// Input: num = 5
// Output: 2
// Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
//
// Example 2:
// Input: num = 1
// Output: 0
// Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
//
// Constraints:
// The given integer num is guaranteed to fit within the range of a 32-bit signed integer.
// num >= 1
// You could assume no leading zero bit in the integer’s binary representation.
// This question is the same as 1009: https://leetcode.com/problems/complement-of-base-10-integer/