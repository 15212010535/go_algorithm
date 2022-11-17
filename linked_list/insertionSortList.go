package main

/*
	对链表进行插入排序
*/

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			// 头指针
			prev := dummyHead
			// 找到需要插入的位置
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// 充当临时变量
			lastSorted.Next = prev.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
