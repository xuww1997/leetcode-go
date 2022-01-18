package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//slow代表中间节点
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	node := &ListNode{0, nil}
	res := node
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			left = left.Next
		} else {
			node.Next = right
			right = right.Next
		}
		node = node.Next
	}
	if left != nil {
		node.Next = left
	} else if right != nil {
		node.Next = right
	}
	return res.Next
}

func main() {
	n4 := &ListNode{3, nil}
	n3 := &ListNode{1, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{4, n2}
	res := sortList(n1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}
