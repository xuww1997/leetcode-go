package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func maxDepth(root *TreeNode) int {
//	if root == nil{
//		return 0
//	}
//	depth := 0
//	var backtrack func(node *TreeNode, d int)
//	backtrack = func(node *TreeNode, d int) {
//		if node.Left == nil && node.Right == nil {
//			if d > depth {
//				depth = d
//			}
//			return
//		}
//		if nextNode := node.Left; nextNode != nil {
//			backtrack(nextNode, d+1)
//		}
//		if nextNode := node.Right; nextNode != nil {
//			backtrack(nextNode, d+1)
//		}
//	}
//	backtrack(root, 1)
//	return depth
//}

//func maxDepth(root *TreeNode) int {
//	if root == nil{
//		return 0
//	}
//	return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
//}
//func max(a,b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	newQueue := make([]*TreeNode, 0)
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			if node := queue[i].Left; node != nil {
				newQueue = append(newQueue, node)
			}
			if node := queue[i].Right; node != nil {
				newQueue = append(newQueue, node)
			}
		}
		depth++
		queue = newQueue
		newQueue = make([]*TreeNode, 0)
	}
	return depth
}

func main() {

}
