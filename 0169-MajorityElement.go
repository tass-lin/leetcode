package main

func majorityElement(nums []int) int {
	var index int
	size := len(nums)
	m := make(map[int]int,size)
	temp := m[nums[0]]

	for i := 0; i < size; i++ {
		m[nums[i]] += 1
		if temp < m[nums[i]] {
			temp =  m[nums[i]]
			index = i
		}
	}

	return nums[index]
}
func majorityElement1(nums []int) int {
	size := len(nums)
	m := make(map[int]int,size)
	for i := 0; i < size; i++ {
		m[nums[i]] += 1
		if m[nums[i]] > size / 2 {
			return nums[i]
		}
	}
	return nums[0]
}

func main()  {
	println(majorityElement([]int{2,2,1,1,1,2,2}))
	println(majorityElement([]int{2,2,3}))

	println(majorityElement1([]int{2,2,1,1,1,2,2}))
	println(majorityElement1([]int{2,2,3}))
}

//Example 1:
//
//Input: [3,2,3]
//Output: 3

//Example 2:
//
//Input: [2,2,1,1,1,2,2]
//Output: 2