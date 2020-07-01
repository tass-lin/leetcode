package main

func validPalindrome(s string) bool {
	var temp int
	i, j := 0, len(s) - 1

	for i < j {
		if s[i] != s[j] {
			temp++
			if s[i+1] == s[j] {
				i++
			} else if s[i] == s[j-1] {
				j--
			} else {
				temp++
			}
		}
		i++
		j--
	}
	var temp2 int
	i, j = 0, len(s) - 1

	for i < j {
		if s[i] != s[j] {
			temp2++
			if s[i] == s[j-1] {
				j--
			} else if s[i+1] == s[j] {
				i++
			} else {
				temp2++
			}
		}
		i++
		j--
	}

	return temp <= 1 || temp2 <= 1
}

func validPalindrome1(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return check(s[i+1:j+1]) || check(s[i:j])
		}
		i++
		j--
	}

	return true
}
func check(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
/*
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True

Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.

Note:
The string will only contain lowercase characters a-z. The maximum length of the string is 50000.
*/