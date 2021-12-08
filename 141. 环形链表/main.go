package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	first, second := head.Next, head
	for first != second {
		if first == nil || first.Next == nil {
			return false
		}
		first = first.Next.Next
		second = second.Next
	}
	return true
}
