package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nil || nodeB != nil {
		if nodeA == nodeB {
			return nodeA
		}
		if nodeA == nil {
			nodeA = headB
			nodeB = nodeB.Next
		} else if nodeB == nil {
			nodeB = headA
			nodeA = nodeA.Next
		} else {
			nodeA = nodeA.Next
			nodeB = nodeB.Next
		}
	}
	return nil
}

func main() {
	c3 := &ListNode{33, nil}
	c2 := &ListNode{32, c3}
	c1 := &ListNode{31, c2}
	a2 := &ListNode{12, c1}
	a1 := &ListNode{11, a2}
	b3 := &ListNode{23, c1}
	b2 := &ListNode{22, b3}
	b1 := &ListNode{21, b2}
	result := getIntersectionNode(a1, b1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
