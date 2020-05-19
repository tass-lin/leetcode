package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	countDash := strings.Count(S,"-")
	count := len(S) - countDash

	if count == 0 {
		return ""
	}

	S = strings.Replace(S,"-","",-1) // -1 全部替換
	S = strings.ToUpper(S) // 轉換大寫

	// (count+K-1)/K - 1 為 (-) 所需的數量
	res := make([]byte, (count+K-1)/K - 1 + count )

	resLen, SLen := len(res), len(S)

	// 從尾開始
	for SLen - K >= 0 {
		copy(res[resLen-K:resLen],S[SLen-K:SLen])
		if SLen - K - 1 >= 0 {
			res[resLen-K-1] = '-'
		}
		resLen -= K + 1
		SLen -= K

	}

	if SLen > 0 {
		copy(res[:resLen],S[:SLen])
	}
	return string(res)
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w",4)) // "5F3Z-2E9W"
	fmt.Println(licenseKeyFormatting("2-5g-3-J",2)) // "2-5G-3J""
}

// You are given a license key represented as a string S which consists only alphanumeric character and dashes.
// The string is separated into N+1 groups by N dashes.
// Given a number K, we would want to reformat the strings such that each group contains exactly K characters,
// except for the first group which could be shorter than K, but still must contain at least one character.
//Furthermore, there must be a dash inserted between two groups and all lowercase letters should be converted to uppercase.
// Given a non-empty string S and a number K, format the string according to the rules described above.
//
// Example 1:
// Input: S = "5F3Z-2e-9-w", K = 4
// Output: "5F3Z-2E9W"
// Explanation: The string S has been split into two parts, each part has 4 characters.
// Note that the two extra dashes are not needed and can be removed.
//
// Example 2:
// Input: S = "2-5g-3-J", K = 2
// Output: "2-5G-3J"
// Explanation: The string S has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.
// Note:
//
// The length of string S will not exceed 12,000, and K is a positive integer.
// String S consists only of alphanumerical characters (a-z and/or A-Z and/or 0-9) and dashes(-).
// String S is non-empty.