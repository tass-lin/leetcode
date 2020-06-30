package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	patternArr := strings.Split(pattern,"")
	strArr := strings.Split(str," ")
	if len(patternArr) != len(strArr) {
		return false
	}
	//if match(patternArr) == false || match(strArr) == false {
	//	return false
	//}
	match(patternArr, strArr)
	return true

}
func match(arr1, arr2 []string) bool{

	return true
}

func main() {
	fmt.Println(wordPattern("abba","dog cat cat dog")) //true
	fmt.Println(wordPattern("abba","dog cat cat fish")) //false
	fmt.Println(wordPattern("aaaa","dog cat cat dog")) //false
	fmt.Println(wordPattern("abba","dog dog dog dog")) //false
	fmt.Println(wordPattern("aba","cat cat cat dog")) //false
	fmt.Println(wordPattern("abaaba","dog cat fish fish cat dog")) //false
}

//Given a pattern and a string str, find if str follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.
//
//Example 1:
//
//Input: pattern = "abba", str = "dog cat cat dog"
//Output: true

//Example 2:
//
//Input:pattern = "abba", str = "dog cat cat fish"
//Output: false

//Example 3:
//
//Input: pattern = "aaaa", str = "dog cat cat dog"
//Output: false

//Example 4:
//
//Input: pattern = "abba", str = "dog dog dog dog"
//Output: false
//Notes:
//You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.