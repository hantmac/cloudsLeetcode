package main

func deleteDuplicates(head *LinkNode) *LinkNode {

	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}

	return head
}
