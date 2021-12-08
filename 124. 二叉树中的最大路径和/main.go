package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		maxPath = max(maxPath, left+right+node.Val)
		return max(left, right) + node.Val
	}
	dfs(root)
	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
