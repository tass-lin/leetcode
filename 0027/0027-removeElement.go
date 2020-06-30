package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1

	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Printf("%v",nums)
	return i
}

func main(){
	//nums := []int{3,2,2,3} // len = 4
	nums := []int{0,1,2,2,3,0,4,2} // len = 8

	println(removeElement(nums,2))
}