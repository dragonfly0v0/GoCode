class Solution(object):
    def minSubArrayLen(self, target, nums):
        """
        :type target: int
        :type nums: List[int]
        :rtype: int
        """
        if nums is None:
            return 0

        i = 0
        res = float("inf")
        sum = 0
        for j in range(len(nums)):
            sum += nums[j]

            while(sum >= target):
                sublength = j - i + 1
                res = min(sublength, res)

                sum -= nums[i]
                i += 1
        print(res)
        return 0 if res == float("inf") else res
    
Solution().minSubArrayLen(7, [2, 3, 1, 2, 4, 3])