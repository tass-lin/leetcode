package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m, max, res := make(map[int]int), -1, []int{}
	search(root, m)

	for index, element := range m {
		if max <= element {
			if max < element {
				max = element
				res = res[0:0]
			}
			res = append(res,index)
		}
	}
	return res
}

func search(root *TreeNode, m map[int]int)  {
	if root == nil{
		return
	}

	m[root.Val]++

	search(root.Left, m)
	search(root.Right, m)
}