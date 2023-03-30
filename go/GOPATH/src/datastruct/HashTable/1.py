from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        if len(nums) == 0 :
            return None
	    
        val_dict = {}
        for i in range(len(nums)):
            s = target - nums[i]

            if s in val_dict:
                return [i, val_dict[s]]
            
            val_dict[nums[i]] = i

        return None