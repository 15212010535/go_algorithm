package main

/*
	链表的中间结点
*/

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	// 判断奇偶
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length%2 == 0 {
		return p1.Next
	}
	return p1
}
