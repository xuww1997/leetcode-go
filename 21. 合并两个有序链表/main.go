package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	//当前节点，即尾节点
	nowNode := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			nowNode.Next = l1
			l1 = l1.Next
		} else {
			nowNode.Next = l2
			l2 = l2.Next
		}
		nowNode = nowNode.Next
	}
	//将l1或l2不为nil的部分添加在尾节点后边
	if l1 != nil {
		nowNode.Next = l1
	} else if l2 != nil {
		nowNode.Next = l2
	}
	//一开始定义的头节点在真实的头节点前一位，因此返回他的next
	return result.Next
}
