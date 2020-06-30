package main

type Node struct {
	Val int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	
	m := 1

	for _, element := range root.Children {
		m = max(m, maxDepth(element) + 1)
	}
	return m
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}