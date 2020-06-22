package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	m2 := make(map[string]int, len(list2))
	var res []string
	temp := math.MaxInt32

	for index, element := range list2 {
		m2[element] = index
	}

	for i, element := range list1 {
		if j, ok := m2[element]; ok {
			if temp == i + j {
				res = append(res, element)
			}
			if temp > i + j {
				temp = i + j
				res = append(res[len(res):], element)
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n", findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"})) // ["Shogun"]
	fmt.Printf("%v\n", findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"})) // ["Shogun"]
	fmt.Printf("%v\n", findRestaurant([]string{"Shogun","Tapioca Express","Burger King","KFC"}, []string{"KFC","Burger King","Tapioca Express","Shogun"})) // ["KFC","Burger King","Tapioca Express","Shogun"]
}

// Suppose Andy and Doris want to choose a restaurant for dinner,
// and they both have a list of favorite restaurants represented by strings.
//
// You need to help them find out their common interest with the least list index sum.
// If there is a choice tie between answers, output all of them with no order requirement.
// You could assume there always exists an answer.
//
// Example 1:
// Input:
// ["Shogun", "Tapioca Express", "Burger King", "KFC"]
// ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
// Output: ["Shogun"]
// Explanation: The only restaurant they both like is "Shogun".
//
// Example 2:
// Input:
// ["Shogun", "Tapioca Express", "Burger King", "KFC"]
// ["KFC", "Shogun", "Burger King"]
// Output: ["Shogun"]
// Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
//
// Note:
// The length of both lists will be in the range of [1, 1000].
// The length of strings in both lists will be in the range of [1, 30].
// The index is starting from 0 to the list length minus 1.
// No duplicates in both lists.