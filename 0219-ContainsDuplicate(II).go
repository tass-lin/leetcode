package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 1 {
		return false
	}
	m := make(map[int]int,len(nums))

	for index, element := range nums {
		if m[element] != 0 && index - (m[element] - 1) <= k  {
			return true
		}
		m[element] = index + 1
	}
	return false
}

func main() {
	println(containsNearbyDuplicate([]int{1,2,3,1},3)) //true
	println(containsNearbyDuplicate([]int{1,0,1,1},1)) //true
	println(containsNearbyDuplicate([]int{1,2,3,1,2,3},2)) //false
	println(containsNearbyDuplicate([]int{99,99},2)) //true
}

//Example 1:
//
//Input: nums = [1,2,3,1], k = 3
//Output: true

//Example 2:
//
//Input: nums = [1,0,1,1], k = 1
//Output: true

//Example 3:
//
//Input: nums = [1,2,3,1,2,3], k = 2
//Output: false