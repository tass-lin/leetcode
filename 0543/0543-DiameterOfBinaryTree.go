package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := goTree(root)
	return res
}

func goTree(root *TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	leftLen, leftDiameter := goTree(root.Left)
	rightLen, rightDiameter := goTree(root.Right)

	length = max(leftLen, rightLen) + 1
	diameter = max(leftLen + rightLen, max(leftDiameter, rightDiameter))
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Given a binary tree, you need to compute the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.
//
// Example:
// Given a binary tree
//      1
//     / \
//    2   3
//   / \
//  4   5
// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
//
// Note: The length of path between two nodes is represented by the number of edges between them.