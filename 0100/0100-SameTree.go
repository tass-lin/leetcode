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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	
	if p.Val != q.Val {
		return false
	}
	
	if isSameTree(p.Left,q.Left) == false {
		return false
	}
	if isSameTree(p.Right,q.Right) == false {
		return false
	}
	
	return true
}

func main() {
	// [1,2,3],[1,2,3]
	treeNode := new(TreeNode)
	treeNodeLeft := new(TreeNode)
	treeNodeRight := new(TreeNode)

	treeNodeLeft.Val = 2
	treeNodeRight.Val = 3

	treeNode.Val = 1
	treeNode.Left = treeNodeLeft
	treeNode.Right = treeNodeRight

	println(isSameTree(treeNode,treeNode))

	// [1,2],[1,null]
	treeNode1 := new(TreeNode)
	treeNodeLeft1 := new(TreeNode)

	treeNodeLeft1.Val = 2

	treeNode1.Val = 1
	treeNode1.Left = treeNodeLeft1

	treeNode2 := new(TreeNode)
	treeNodeLeft2 := new(TreeNode)
	treeNodeRight2 := new(TreeNode)

	//treeNodeLeft2.Val = 2
	treeNodeRight2.Val = 2

	treeNode2.Val = 1
	treeNode2.Left = treeNodeLeft2

	println(isSameTree(treeNode1,treeNode2))

}
