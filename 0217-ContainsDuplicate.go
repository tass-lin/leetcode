package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int,len(nums))

	for _, element := range nums {
		if m[element] == 1{
			return true
		}
		m[element] += 1
	}
	return false
}
func containsDuplicate1(nums []int) bool {
	m := make(map[int]bool,len(nums))

	for _, element := range nums {
		if m[element] {
			return true
		}
		m[element] = true
	}
	return false
}

func main() {
	println(containsDuplicate([]int{1,2,3,1}))
	println(containsDuplicate1([]int{1,2,3,1}))
}

//Example 1:
//
//Input: [1,2,3,1]
//Output: true

//Example 2:
//
//Input: [1,2,3,4]
//Output: false

//Example 3:
//
//Input: [1,1,1,3,3,4,3,2,4,2]
//Output: true