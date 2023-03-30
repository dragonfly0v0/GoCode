from typing import Optional

class ListNode():
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        if head is None:
            return None
        
        dummy = ListNode(next=head)
        slow = dummy
        fast = head
        i = 1

        while fast != None :
            fast = fast.next

            if i > n :
                slow = slow.next
            
            i += 1
        
        slow.next = slow.next.next

        return dummy.next