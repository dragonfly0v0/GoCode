from typing import Optional

class ListNode():
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return None
        
        fast = head
        slow = head

        while fast != None and fast.next != None:
            fast = fast.next.next
            slow = slow.next

            if fast == slow :
                index1 = fast
                index2 = head

                while index1 != index2:
                    index1 = index1.next
                    index2 = index2.next
            return index2
        return None