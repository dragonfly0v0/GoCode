package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var pre *ListNode
	var temp *ListNode
	pre = nil

	for cur != nil {
		temp = cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
