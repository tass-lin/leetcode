package main

import "fmt"

func isSubsequence(s string, t string) bool {
	temp1, temp2 := []byte(s), []byte(t)
	i, sSize := 0, len(temp1)
	if sSize == 0 {
		return true
	}
	for _, element := range temp2 {
		if element == temp1[i] {
			i++
			if i == sSize {
				return true
			}
		}
	}
	return false
}
func isSubsequence1(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j]{
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc","ahbgdc")) // true
	fmt.Println(isSubsequence("axc","ahbgdc")) // false
	fmt.Println(isSubsequence("","ahbgdc")) // true

	fmt.Println(isSubsequence1("abc","ahbgdc")) // true
	fmt.Println(isSubsequence1("axc","ahbgdc")) // false
	fmt.Println(isSubsequence1("","ahbgdc")) // true
}

// Given a string s and a string t, check if s is subsequence of t.
// A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).
//
// Follow up:
// If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?
//
// Credits:
// Special thanks to @pbrother for adding this problem and creating all test cases.
//
// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true
//
// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false
//
// Constraints:
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// Both strings consists only of lowercase characters.