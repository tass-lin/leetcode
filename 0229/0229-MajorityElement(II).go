package main

import "fmt"

func majorityElement(nums []int) []int {
	var res []int
	size := len(nums)
	m := make(map[int]int, size)

	for i := 0; i < size; i++ {
		m[nums[i]] += 1
		if  m[nums[i]] > size / 3 {
			res = append(res,nums[i])
			m[nums[i]] = -1
		}
	}
	return res
}


func main()  {
	fmt.Printf("%v",majorityElement([]int{3,2,3})) // [3]
	fmt.Printf("%v",majorityElement([]int{1,1,1,3,3,2,2,2})) //[1,2]
	fmt.Printf("%v",majorityElement([]int{2,2})) //[2]
}
