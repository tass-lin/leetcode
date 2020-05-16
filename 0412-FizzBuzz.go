package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		switch {
		case i % 15 == 0:
			res = append(res,"FizzBuzz")
		case i % 3 == 0:
			res = append(res,"Fizz")
		case i % 5 == 0:
			res = append(res,"Buzz")
		default:
			res = append(res,strconv.Itoa(i))
		}
	}
	return res
}

func main() {
	fmt.Printf("%v",fizzBuzz(15))
}

// Write a program that outputs the string representation of numbers from 1 to n.
// But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”.
// For numbers which are multiples of both three and five output “FizzBuzz”.
// 
// Example:
// n = 15,
// Return:
// [
// "1",
// "2",
// "Fizz",
// "4",
// "Buzz",
// "Fizz",
// "7",
// "8",
// "Fizz",
// "Buzz",
// "11",
// "Fizz",
// "13",
// "14",
// "FizzBuzz"
// ]