package main

import "fmt"

var s []byte
func reverseString(s []byte)  {
	min, max := 0, len(s)-1
	var temp byte
	for min < max {
		temp = s[max]
		s[max] = s[min]
		s[min] = temp
		min++
		max--
	}
}

func reverseString1(s []byte)  {
	min, max := 0, len(s)-1
	for min < max {
		s[min], s[max] = s[max], s[min]
		min++
		max--
	}
}

func main() {
	str := []byte("hello")
	s = str
	reverseString(s)
	fmt.Printf("%v\n",string(s)) // olleh

	str1 := []byte("Hannah")
	s = str1
	reverseString(s)
	fmt.Printf("%v\n",string(s)) // hannaH


	str2 := []byte("hello")
	s = str2
	reverseString1(s)
	fmt.Printf("%v\n",string(s)) // olleh

	str3 := []byte("Hannah")
	s = str3
	reverseString1(s)
	fmt.Printf("%v",string(s)) // hannaH
}

// Write a function that reverses a string. The input string is given as an array of characters char[].
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// You may assume all the characters consist of printable ascii characters.
//
// Example 1:
// Input: ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
//
// Example 2:
// Input: ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]