package main

import "fmt"

func removeDuplicates(nums []int) int {
	left,right,size := 0,1,len(nums)

	for ;right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left],nums[right] = nums[right],nums[left]
		fmt.Printf("%v",nums)
	}
	return left+1
}

func main(){
	nums := []int{0,0,0,1,2,3,3,4,4,4}
	//nums := []int{0,1,2,2,3,4}
	println(removeDuplicates(nums))
}