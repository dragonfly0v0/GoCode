class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        if len(s) == 0:
            return ""
        length = len(s)
        str = list(s)
        for i in range(0, length, 2*k):
            if i + k < length:
                substr = str[i:i+k]
                self.reverse(substr)
            else:
                substr = str[i:length]
                self.reverse(substr)
            s = s[:i] + ''.join(substr) + s[i+k:]
        return s


    def reverse(self, s):
        i = 0
        j = len(s) - 1
        while i < j:
            s[i], s[j] = s[j], s[i]
            i += 1
            j -= 1

s = "abgdhjdeui"
res = Solution().reverseStr(s, 3)
print(res)