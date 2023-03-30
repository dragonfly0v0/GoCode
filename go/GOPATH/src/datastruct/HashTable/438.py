from typing import List
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if s == "" or p == "":
            return None
        from collections import defaultdict
        freqP = defaultdict(int)
        for ch in p:
            freqP[ord(ch)] += 1
        print("frepP={}\n".format(freqP))

        res = []
        freqS = defaultdict(int)
        left, right = 0, 0

        for ch in s:
            freqS[ord(ch)] += 1
            right += 1

            print("frepS={}\n".format(freqS))

            while freqS[ord(ch)] > freqP[ord(ch)]:
                print("freqS[ch] > freqP[ch], frepS={}\n".format(freqS))
                freqS[ord(s[left])] -= 1
                left += 1

            if right - left == len(p):
                print("res = {}, left={}".format(res, left))
                res.append(left)

        return res

    
s = "cbaebabacd"
p = "abc"
res = Solution().findAnagrams(s, p)
print(res)
        