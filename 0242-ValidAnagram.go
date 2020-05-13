package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr, tArr := []rune(s), []rune(t)
	temp := make(map[rune]int,len(sArr))

	for index := range sArr {
		temp[sArr[index]]++
		temp[tArr[index]]--
	}
	for _, element := range temp {
		if element != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram","nagaram"))
	fmt.Println(isAnagram("rat","car"))
}

//重複字母組成另外一個單字


//Given two strings s and t , write a function to determine if t is an anagram of s.
//
//Example 1:
//
//Input: s = "anagram", t = "nagaram"
//Output: true
//Example 2:
//
//Input: s = "rat", t = "car"
//Output: false
//Note:
//You may assume the string contains only lowercase alphabets.
//
//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?
