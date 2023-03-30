package main

type ListNode struct {
	Data int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var slow *ListNode
	fast := head
	for head != nil {
		fast = head.Next
		head.Next = slow
		slow = head
		head = fast
	}
	return head
}
