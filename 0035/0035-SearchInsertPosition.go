package main

func searchInsert(nums []int, target int) int {
	for index, element := range nums {
		if target-element == 0{
			return index
		} else if target-element < 0 {
			return index
		}
	}
	return len(nums)
}

func searchInsert1(nums []int, target int) int {
	// 没有把i放入for语句中
	// 是为了兼容，len(nums) == 0 和 target > nums[len(nums)-1]两种情况
	i := 0

	for i < len(nums) && nums[i] <= target {
		// 相等的时候，直接返回
		if nums[i] == target {
			return i
		}

		// 否则，就去检查下一个
		i++
	}

	return i
}

func main()  {
	println(searchInsert1([]int{1,3,5,6},5)) //2
	println(searchInsert1([]int{1,3,5,6},2)) //1
	println(searchInsert1([]int{1,3,5,6},7)) //4
	println(searchInsert1([]int{1,3,5,6},0)) //0
}

