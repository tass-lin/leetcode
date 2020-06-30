package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	field := strings.Fields(s)
	return len(field)
}

func main() {
	fmt.Println(countSegments("Hello, my name is John")) // 5
}

// Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.
// Please note that the string does not contain any non-printable characters.
//
// Example:
// Input: "Hello, my name is John"
// Output: 5