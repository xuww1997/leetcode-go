package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nowNode := head
	num := 1
	for nowNode.Next != nil {
		num++
		nowNode = nowNode.Next
	}
	k = k % num
	if k == 0 {
		return head
	}
	//此时nowNode指尾节点，首位相连
	nowNode.Next = head
	for i := 0; i < num-k; i++ {
		nowNode = nowNode.Next
	}
	//此时nowNode为新
	result := nowNode.Next
	nowNode.Next = nil
	return result
}

func main() {
	//node5 := &ListNode{5,nil}
	//node4 := &ListNode{4,node5}
	//node3 := &ListNode{3,node4}
	//node2 := &ListNode{2,node3}
	//node1 := &ListNode{1,node2}
	node3 := &ListNode{2, nil}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{0, node2}
	result := rotateRight(node1, 4)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
