package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = res
		res = head
		head = nextNode
	}
	return res
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	res := reverseList(node1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
