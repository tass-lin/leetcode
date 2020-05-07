package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode

	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return head
}

func main(){
	listNode1 := new(ListNode)
	listNode1.Val = 1

	listNode2 := new(ListNode)
	listNode2.Val = 2

	listNode3 := new(ListNode)
	listNode3.Val = 4

	listNode4 := new(ListNode)
	listNode4.Val = 1

	listNode5 := new(ListNode)
	listNode5.Val = 3

	listNode6 := new(ListNode)
	listNode6.Val = 4

	listNode1.Next = listNode2
	listNode2.Next = listNode3

	listNode4.Next = listNode5
	listNode5.Next = listNode6

	fmt.Printf("%+v\n", listNode1)

	fmt.Printf("%+v",mergeTwoLists(listNode1,listNode4))
}
