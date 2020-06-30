package main

import "fmt"

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := head

	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return  head
}

func main()  {
	//[1,1,2,3,3]

	listNode1 := new(ListNode)
	listNode1.Val = 1

	listNode2 := new(ListNode)
	listNode2.Val = 1

	listNode3 := new(ListNode)
	listNode3.Val = 2

	listNode1.Next = listNode2
	listNode2.Next = listNode3

	fmt.Printf("%+v\n", deleteDuplicates(listNode1))
}