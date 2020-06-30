package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for index, element := range numbers {
		if m[target-element] != 0 {
			return []int{m[target-element], index + 1}
		}
		m[element] = index + 1
	}
	return nil
}

func twoSum1(numbers []int, target int) []int {
	var res []int
	size := len(numbers)

	for i, j := 0, size - 1; i < j; {
		if target == numbers[i] + numbers[j] {
			res = []int{i+1, j+1}
			return res
		} else if target > numbers[i] + numbers[j] {
			i++
		} else if target < numbers[i] + numbers[j] {
			j--
		}
	}

	return nil
}

func main()  {
	fmt.Printf("%v",twoSum([]int{2,7,11,15},26))
	fmt.Printf("%v",twoSum1([]int{2,7,11,15},26))
}

//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.