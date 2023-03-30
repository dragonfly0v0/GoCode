package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var dummy_head *ListNode
	dummy_head.Next = head
	cur := dummy_head

	var temp *ListNode

	for cur.Next != nil && cur.Next.Next != nil {
		temp = head.Next.Next
		cur.Next = head.Next
		head.Next.Next = head
		head.Next = temp

		cur = head
		head = temp
	}
	return dummy_head.Next
}
