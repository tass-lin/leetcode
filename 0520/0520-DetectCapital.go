package main

import "fmt"

func detectCapitalUse(word string) bool {
	firstWord, tailWord := word[0], word[1:]
	switch {
	case firstWord <= 90 :
		for _, element := range tailWord {
			if element <= 90 {
				return uppercase(tailWord)
			} else {
				return lowercase(tailWord)
			}
		}
	default:
		return lowercase(tailWord)
	}
	return true
}

func uppercase(tailWord string) bool {
	for _, element := range tailWord {
		if element >= 97 {
			return false
		}
	}
	return true
}

func lowercase(tailWord string) bool  {
	for _, element := range tailWord {
		if element <= 90 {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(byte('A'))
	fmt.Println(byte('Z'))
	fmt.Println(byte('a'))
	fmt.Println(byte('z'))
	fmt.Println(detectCapitalUse("USA")) // True
	fmt.Println(detectCapitalUse("FlaG")) // False
	fmt.Println(detectCapitalUse("uZfa")) // False
	fmt.Println(detectCapitalUse("FFFFFFFFFFFFFFFFFFFFf")) // False
}

// Given a word, you need to judge whether the usage of capitals in it is right or not.
// We define the usage of capitals in a word to be right when one of the following cases holds:
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Otherwise, we define that this word doesn't use capitals in a right way.
//
// Example 1:
// Input: "USA"
// Output: True
//
// Example 2:
// Input: "FlaG"
// Output: False