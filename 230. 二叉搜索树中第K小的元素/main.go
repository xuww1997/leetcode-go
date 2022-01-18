package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	sortedArr := make([]int, 0)
	var backtrack func(node *TreeNode)
	backtrack = func(node *TreeNode) {
		if node == nil {
			return
		}
		backtrack(node.Left)
		sortedArr = append(sortedArr, node.Val)
		backtrack(node.Right)
	}
	backtrack(root)
	fmt.Println(sortedArr)
	return sortedArr[k-1]
}

func main() {
	node2 := &TreeNode{2, nil, nil}
	node1 := &TreeNode{1, nil, node2}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, node1, node4}
	kthSmallest(node3, 1)
}
