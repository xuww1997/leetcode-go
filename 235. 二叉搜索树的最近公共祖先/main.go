package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	var result *TreeNode
//	pPath := getPath(p, root)
//	qPath := getPath(q, root)
//	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i] == qPath[i]; i++ {
//		result = pPath[i]
//	}
//	return result
//}
//
//func getPath(target *TreeNode, root *TreeNode) (path []*TreeNode) {
//	node := root
//	for target != node {
//		path = append(path, node)
//		if target.Val < node.Val {
//			node = node.Left
//		} else {
//			node = node.Right
//		}
//	}
//	path = append(path, node)
//	return
//}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}
