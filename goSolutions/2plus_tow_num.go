package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PlusTwoNum(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = new(ListNode)
	var curr = dummy
	var carry = 0
	var x = 0
	var y = 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
		}else {
			x=0
		}
		if l2 != nil {
			y = l2.Val
		}else {
			y=0
		}
		s := carry + x + y
		carry = s / 10
		fmt.Println(carry)
		curr.Next = &ListNode{s % 10, nil}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{1, nil}
	}
	//fmt.Println(curr.Next)
	//fmt.Println(curr.Val)
	return dummy.Next
}

func main() {
	var l1 = new(ListNode)
	l1.Val = 9
	l1.Next = &ListNode{8,nil}
	var l2 = new(ListNode)
	l2.Val = 1
	res := PlusTwoNum(l1, l2)
	fmt.Println(res)
}
