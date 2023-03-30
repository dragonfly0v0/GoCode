from typing import List
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if len(haystack) == 0 or len(needle) == 0:
            return -1
        
        def GetNext(neddle: str):
            j = 0  # j是前缀末尾，也代表i包括i在内的最大相同前后缀的长度
            
            next = [0]*len(needle)
            neddle = list(neddle)
            for i in range(1, len(neddle)):
                while j > 0 and neddle[i] != neddle[j]:
                    j = next[j - 1]
                if neddle[i] == neddle[j]:
                    j += 1
                next[i] = j

            return next
        
        # needle = "ababab"
        # string = list(needle)
        # print(GetNext(next, string))
       
        next = GetNext(needle)
        print(next)

        next2 = self.sgetNext(needle)
        print(next2)

        hs = haystack
  
        j = 0
        for i in range(0, len(haystack)):
            while j > 0 and hs[i] != needle[j]:
                j = next[j-1]
            if hs[i] == needle[j]:
                j += 1
            if j == len(next):
                return i - len(next) + 1
        return -1

    def sgetNext(self, needle):
        next = [0] * len(needle)
        j = 0
        next[0] = j
        for i in range(1, len(needle)):
            while j >= 1 and needle[i] != needle[j]:
                j = next[j-1]
            if needle[i] == needle[j]:
                j += 1
            next[i] = j
        return next

# class Solution:
#     def strStr(self, haystack: str, needle: str) -> int:
#         if len(needle) == 0:
#             return 0
#         next = self.getNext(needle)
#         j = 0
#         for i in range(len(haystack)):
#             while j >= 1 and haystack[i] != needle[j]:
#                 j = next[j-1]
#             if haystack[i] == needle[j]:
#                 j += 1
#             if j == len(needle):
#                 return i - len(needle) + 1
#         return -1
    
#     def getNext(self, needle):
#         next = [0] * len(needle)
#         j = 0
#         next[0] = j
#         for i in range(1, len(needle)):
#             while j >= 1 and needle[i] != needle[j]:
#                 j = next[j-1]
#             if needle[i] == needle[j]:
#                 j += 1
#             next[i] = j
#         return next


Str = "casadbutsad"
s = "sad"
print(Solution().strStr(Str, s))