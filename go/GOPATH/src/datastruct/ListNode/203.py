# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __len__(self):
         cur = self
         count = 0
         while cur != None:
              count += 1
              cur = cur.next
         return count


class Solution(object):
    def removeElements(self, head: ListNode, val)->ListNode:
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        dummy_head = ListNode(next=head)
        cur = dummy_head

        while cur.next != None :
                if cur.next.val == val :
                    cur.next = cur.next.next
                else :
                     cur = cur.next
        return dummy_head.next
    
val = 6

head = [1,2,6,3,4,5,6]
head_node = ListNode(head[0])
cur = head_node
for i in range(1, len(head)):
    cur.next = ListNode(head[i])
    cur = cur.next
cur = head_node
res = Solution().removeElements(cur, val)
cur1 = res

while cur1 != None :
    print(cur1.val)
    cur1 = cur1.next