class Solution(object):
    def totalFruit(self, fruits):
        """
        :type fruits: List[int]
        :rtype: int
        """
        from collections import Counter

        if fruits is None :
            return 0


        cnt = Counter()
        left = ans = 0
        for right, x in enumerate(fruits):
            cnt[x] += 1
            while len(cnt) > 2:
                cnt[fruits[left]] -= 1

                if cnt[fruits[left]] == 0:
                    cnt.pop(fruits[left])
                
                left += 1
            ans = max(ans, right - left + 1)
        
        return ans

