package main

import "fmt"

func maxPower(s string) int {
	temp, max, maxTemp := s[0], 0, 0

	for i := 0; i < len(s); i++ {
		if temp == s[i] {
			maxTemp++
		} else {
			temp = s[i]
			maxTemp = 1
		}
		if max < maxTemp{
			max = maxTemp
		}
	}
	return max
}

func main() {
	fmt.Println(maxPower("leetcode")) // 2 ee
	fmt.Println(maxPower("abbcccddddeeeeedcba")) // 5 eeeee
	fmt.Println(maxPower("triplepillooooow")) // 5
	fmt.Println(maxPower("hooraaaaaaaaaaay")) // 11
	fmt.Println(maxPower("tourist")) // 1
	fmt.Println(maxPower("b")) // 1
	fmt.Println(maxPower("cc")) // 2
}

// Given a string s, the power of the string is the maximum length of a non-empty substring that contains only one unique character.
// Return the power of the string.
//
// Example 1:
// Input: s = "leetcode"
// Output: 2
// Explanation: The substring "ee" is of length 2 with the character 'e' only.
//
// Example 2:
// Input: s = "abbcccddddeeeeedcba"
// Output: 5
// Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
//
// Example 3:
// Input: s = "triplepillooooow"
// Output: 5
//
// Example 4:
// Input: s = "hooraaaaaaaaaaay"
// Output: 11
//
// Example 5:
// Input: s = "tourist"
// Output: 1
//
// Constraints:
// 1 <= s.length <= 500
// s contains only lowercase English letters.