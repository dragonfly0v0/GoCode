from typing import List
from collections import defaultdict

class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        if len(nums1) == 0 or len(nums2) == 0 or len(nums3) == 0 or len(nums4) == 0:
            return 0
		    
        
        comdict = defaultdict(int)

        for n1 in nums1:
            for n2 in nums2:
                comdict[n1+n2] += 1

        count = 0

        for n3 in nums3:
            for n4 in nums4:
                count += comdict[-n3-n4]
        return count