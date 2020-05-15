package main

import "fmt"

func firstUniqChar(s string) int {
	m := make(map[int32]int, len(s))
	for _, element  := range s {
		m[element] += 1
	}
	for index, element  := range s {
		if m[element] == 1 {
			return index
		}
	}
	return -1
}


func main() {
	fmt.Println(firstUniqChar("leetcode")) // 0
	fmt.Println(firstUniqChar("loveleetcode")) // 2
}

// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

// Examples:
// s = "leetcode"
// return 0.
//
// s = "loveleetcode",
// return 2.
//
// Note: You may assume the string contain only lowercase letters.