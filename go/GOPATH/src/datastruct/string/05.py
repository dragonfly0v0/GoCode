class Solution:
    def replaceSpace(self, s: str) -> str:
        if len(s) == 0:
            return ""
        
        count = 0 
        str = list(s)
        for i in str:
            if i == " ":
                count += 1
        
        