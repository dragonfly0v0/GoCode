from typing import List

class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) == 0 or len(nums2) == 0:
            return []
        
        val_dict = {}
        for num in nums1:
            val_dict[num] = 1

        res = []
        for num in nums2:
            if num in val_dict.keys() and val_dict[num] == 1:
                res.append(num)
                val_dict[num] = 0

        return res
    

num1 = [1, 2,2,45,4,3,2,7]
num2 = [1, 2, 7]
res = Solution().intersection(num1, num2)
print(res)
