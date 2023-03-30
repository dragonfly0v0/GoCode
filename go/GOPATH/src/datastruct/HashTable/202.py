from typing import List
class Solution:
    def isHappy(self, n: int) -> bool:
        if n == 0:
            return True
        
        def caculate_happy(num):
            sum = 0

            while num:
                sum += (num % 10) ** 2
                num = num // 10

            return sum
        
        record = set()
        while True:
            n = caculate_happy(n)
            if n == 1:
                return True
            if n in record:
                return False
            record.add(n)
