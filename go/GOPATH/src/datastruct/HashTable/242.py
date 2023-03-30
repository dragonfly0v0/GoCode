from typing import Optional

class ListNode():
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if s is None or t is None:
            return False
        
        record = [0] * 26
        
        for i in s:
            record[ord(i) - ord('a')] += 1

        for i in t:
            record[ord(i) - ord('a')] -= 1

        for i in range(26):
            if record[i] != 0:
                return False
        return True