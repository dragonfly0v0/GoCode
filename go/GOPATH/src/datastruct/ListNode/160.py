from typing import Optional

class ListNode():
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        lenA, lenB = 0, 0

        curA = headA
        curB = headB

        while curA != None:
            curA = curA.next
            lenA += 1

        while curB != None:
            curB = curB.next
            lenB += 1

        fast, slow = headA, headB
        if lenA > lenB:
            gap = lenA - lenB
            fast, slow = headA, headB
        else :
            gap = lenB - lenA
            fast, slow = headB, headA
        
        while gap > 0:
            fast = fast.next
            gap -= 1

        while fast != slow:
            fast = fast.next
            slow = slow.next
        
        return fast