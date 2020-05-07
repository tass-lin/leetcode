package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res [][]int
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res = [][]int{}
	nextSliceNode := []*TreeNode{root}
	levelAppend(nextSliceNode)

	return res
}

func levelAppend(nodeArr[]*TreeNode) {
	var nextNodeArr []*TreeNode
	var valArr []int

	len := len(nodeArr)

	if 0 == len {
		return
	}

	for i := 0; i < len; i++ {
		value := nodeArr[i]
		if value.Left != nil {
			nextNodeArr = append(nextNodeArr,value.Left)
		}
		if value.Right != nil {
			nextNodeArr = append(nextNodeArr,value.Right)
		}
		valArr = append(valArr,value.Val)
	}
	levelAppend(nextNodeArr)

	res = append(res,valArr)
}

// [3,9,20,null,null,15,7]

//    3
//   / \
//  9  20
//    /  \
//   15   7

//[
//	[15,7],
//	[9,20],
//	[3]
//]