package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	word := strings.Fields(s)
	for index, element := range word {
		bytes := []byte(element)
		i, j := 0, len(bytes)-1
		for i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
		word[index] = string(bytes)
	}

	return strings.Join(word, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest")) // s'teL ekat edoCteeL tsetnoc
}

// Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
//
// Example 1:
// Input: "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Note: In the string, each word is separated by single space and there will not be any extra space in the string.