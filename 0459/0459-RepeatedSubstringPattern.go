package main

import (
	"fmt"
	"strings"
)
func repeatedSubstringPattern(s string) bool {
	size, i := len(s), 1
	if size == 2 {
		return check(s,1)
	}
	for i <= size / 2  {
		if size % i != 0 {
			i++
			continue
		}
		if check(s, i) {
			return true
		}
		i++
	}
	return false
}
func check(s string,max int) bool {
	size, temp := len(s), max
	for i := 1; i < size/max; i++ {
		if strings.Contains(s[(max*i)-temp:(max*i)],s[(max*i):(max*i)+temp]) == false {
			return false
		}
	}
	return true
}

func repeatedSubstringPattern1(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)

	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}


func main() {
	fmt.Println(repeatedSubstringPattern("ab")) // false size := 2
	fmt.Println(repeatedSubstringPattern("bb")) // true size := 2
	fmt.Println(repeatedSubstringPattern("abab")) // true [:2] [2:] size := 4
	fmt.Println(repeatedSubstringPattern("aba")) // False
	fmt.Println(repeatedSubstringPattern("abcabcabcabc")) // true [:6] [6:] size := 12
	fmt.Println(repeatedSubstringPattern("ababab")) // true [:2] [2:4] [4:] size := 6
	fmt.Println(repeatedSubstringPattern("ababba")) // false
	fmt.Println(repeatedSubstringPattern("abaababaab")) // true size := 10
	fmt.Println(repeatedSubstringPattern("ababababab")) // true size := 10
	fmt.Println(repeatedSubstringPattern("babbabbabbabbab")) // true size := 15
}

// Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
// You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.



// Example 1:
// Input: "abab"
// Output: True
// Explanation: It's the substring "ab" twice.
//
// Example 2:
// Input: "aba"
// Output: False
//
// Example 3:
// Input: "abcabcabcabc"
// Output: True
// Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)


// https://leetcode.com/problems/repeated-substring-pattern/discuss/94334/easy-python-solution-with-explaination
// Basic idea:
// First char of input string is first char of repeated substring
// Last char of input string is last char of repeated substring
// Let S1 = S + S (where S in input string)
// Remove 1 and last char of S1. Let this be S2
// If S exists in S2 then return true else false
// Let i be index in S2 where S starts then repeated substring length i + 1 and repeated substring S[0: i+1]