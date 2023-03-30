from typing import Optional

class ListNode():
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


class Solution:
    
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        cur = dummy

        while cur.next != None and cur.next.next != None:
            temp = head.next.next

            cur.next = head.next
            head.next.next = head
            head.next = temp

            cur = head
            head = temp
        
        return dummy.next