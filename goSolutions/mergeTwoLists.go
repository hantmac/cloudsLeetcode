package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

//非递归
//func mergeTwoLists(l1 *LinkNode, l2 *LinkNode) *LinkNode {
//	var headList = &LinkNode{0, nil}
//	var prev = headList
//
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			prev.Next = l1
//			l1 = l1.Next
//		} else {
//			prev.Next = l2
//			l2 = l2.Next
//		}
//		prev = prev.Next
//	}
//
//	if l1 == nil {
//		prev.Next = l2
//	} else {
//		prev.Next = l1
//	}
//
//	return headList.Next
//}

func mergeTwoLists(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1 = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2 = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
