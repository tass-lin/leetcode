package main

import "math"

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return same(root.Left,root.Right)
}

func same(left *TreeNode,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && same(left.Left,right.Right) && same(left.Right,right.Left)
}

//[1,2,2,3,4,4,3]
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3