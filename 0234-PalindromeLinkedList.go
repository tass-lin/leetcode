package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr,head.Val)
		head = head.Next
	}
	size := len(arr)
	for i := 0; i < size / 2; i++ {
		if  arr[i] != arr[size-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr,head.Val)
		head = head.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r]{
			return false
		}
		l++
		r--
	}
	return true
}

func main()  {
	var ListNode1 ListNode
	ListNode1 = ListNode{Val:2,Next:nil}

	var res ListNode
	res = ListNode{Val:1,Next:&ListNode1}

	println(isPalindrome(&res))
	println(isPalindrome1(&res))
}

//Example 1:
//
//Input: 1->2
//Output: false

//Example 2:
//
//Input: 1->2->2->1
//Output: true

//Example 3:
//
//Input: 1->2->3->2->1
//Output: true

