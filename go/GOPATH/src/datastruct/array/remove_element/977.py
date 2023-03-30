## 思路: 
## 1.数组是升序数组，所有最大值肯定在两边
## 2.定义i,j 从两边往中间移动
## 3.定义新的数组去接收它

class Solution(object):
    def sortedSquares(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        ## return sorted(num * num for num in nums)
        left, right, k = 0, len(nums)-1, len(nums)-1
        ans = [-1] * len(nums)
        while (left <= right):
            lm, rm = nums[left]**2, nums[right]**2
            if lm < rm :
                ans[k] = rm
                right -= 1
                
            else:
                ans[k] = lm
                left += 1
            print("ans是{}".format(ans))
            k -= 1
        return ans
    

nums = [-7,-3,2,3,11]
print(Solution().sortedSquares(nums))


