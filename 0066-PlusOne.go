package main

import "fmt"

func plusOne(digits []int) []int {
	dlen := len(digits)

	if dlen == 0 {
		return []int{1}
	}

	digits[dlen-1]++

	for i := dlen - 1; i > 0 ; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main()  {
	fmt.Printf("%v",plusOne([]int{4,3,2,1}))
	fmt.Printf("%v",plusOne([]int{9}))
}