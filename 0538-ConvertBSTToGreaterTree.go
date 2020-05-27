package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	goTree(root, &sum)

	return root
}

func goTree(root *TreeNode, sum *int)  {
	if root == nil {
		return
	}
	goTree(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	goTree(root.Left, sum)
}


// Given a Binary Search Tree (BST),
// convert it to a Greater Tree such that every key of the original BST
// is changed to the original key plus sum of all keys greater than the original key in BST.
//
// Example:
//
// Input: The root of a Binary Search Tree like this:
//    5
//  /   \
// 2     13
//
// Output: The root of a Greater Tree like this:
//     18
//   /   \
// 20     13
// Note: This question is the same as 1038: https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/