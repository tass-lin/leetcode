package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	temp, m, res := make([]int,len(nums)), make(map[int]string,len(nums)), []string{}
	copy(temp,nums)
	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	for index, element := range temp {
		switch {
		case index == 0:
			m[element] = "Gold Medal"
		case index == 1:
			m[element] = "Silver Medal"
		case index == 2:
			m[element] = "Bronze Medal"
		default:
			m[element] = strconv.Itoa(index+1)
		}
	}
	for _, element := range nums {
		res = append(res,m[element])
	}
	return res
}

func main() {
	// fmt.Printf("%v\n",findRelativeRanks([]int{5, 4, 3, 2, 1})) // ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
	fmt.Printf("%v\n",findRelativeRanks([]int{10,3,8,9,4})) // ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
}

// Given scores of N athletes, find their relative ranks and the people with the top three highest scores,
// who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".
//
// Example 1:
// Input: [5, 4, 3, 2, 1]
// Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
// Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
// For the left two athletes, you just need to output their relative ranks according to their scores.
//
// Note:
// N is a positive integer and won't exceed 10,000.
// All the scores of athletes are guaranteed to be unique.