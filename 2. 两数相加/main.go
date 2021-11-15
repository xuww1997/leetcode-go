package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := 0
	var fl = &ListNode{}
	l := fl
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + r
		l.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		r = sum / 10
		l = l.Next
	}
	if r != 0 {
		l.Next = &ListNode{Val: r}
	}

	return fl.Next
}
