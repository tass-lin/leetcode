package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num <= 6 {
		// 6 = 2 × 3
		// 5 = 5
		// 4 = 2 x 2
		// 3 = 3
		// 2 = 2
		// 1 = 1
		return true
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	}

	if num%3 == 0 {
		return isUgly(num / 3)
	}

	if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}

func main() {
	fmt.Println(isUgly(6)) //true
	fmt.Println(isUgly(8)) //true
	fmt.Println(isUgly(14)) //false
}

//Write a program to check whether a given number is an ugly number.
//
//Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

//Example 1:
//Input: 6
//Output: true
//Explanation: 6 = 2 × 3

//Example 2:
//Input: 8
//Output: true
//Explanation: 8 = 2 × 2 × 2

//Example 3:
//Input: 14
//Output: false
//Explanation: 14 is not ugly since it includes another prime factor 7.