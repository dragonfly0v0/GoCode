package main

/*
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	curB := headB

	lenA := 0
	lenB := 0

	for curA != nil {
		curA = curA.Next
		lenA++
	}

	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var fast, slow *ListNode
	var gap int
	if lenA >= lenB {
		gap = lenA - lenB
		fast, slow = headA, headB
	} else {
		gap = lenB - lenA
		fast, slow = headB, headA
	}

	for i := 0; i < gap; i++ {
		fast = fast.Next
		gap--
	}

	for fast != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil

}
