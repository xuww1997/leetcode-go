package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeKLists(lists []*ListNode) *ListNode {
//	head := &ListNode{0, nil}
//	tail := head
//	for true {
//		minLoc := -1
//		for i := 0; i < len(lists); i++ {
//			if lists[i] == nil {
//				continue
//			}
//			if minLoc == -1{
//				minLoc = i
//				continue
//			}
//			if lists[i].Val < lists[minLoc].Val {
//				minLoc = i
//			}
//		}
//		if minLoc == -1{
//			break
//		}
//		minNode := lists[minLoc]
//		tail.Next = minNode
//		lists[minLoc] = minNode.Next
//		tail = tail.Next
//		if len(lists) == 0 {
//			break
//		}
//	}
//	return head.Next
//}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	newLists := make([]*ListNode, 0)
	for i := 0; i < len(lists); i = i + 2 {
		if i+1 <= len(lists)-1 {
			newLists = append(newLists, mergeTwoLists(lists[i], lists[i+1]))
		} else {
			newLists = append(newLists, lists[i])
		}
	}
	return mergeKLists(newLists)
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
