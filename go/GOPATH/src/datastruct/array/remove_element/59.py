from typing import List
class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        startx, starty = 0, 0
        loop = n / 2
        mid = n % 2
        count = 1

        nums = [[0]*n for _ in range(n)]

        for offset in range(1, int(loop)+1):
            for j in range(starty, n - offset):
                nums[startx][j] = count
                count += 1

            for i in range(startx, n - offset):
                nums[i][n - offset] = count
                count += 1

            for j in range(n - offset, starty, -1):
                nums[n - offset][j] = count
                count += 1

            for i in range(n - offset, startx, -1):
                nums[i][starty] = count
                count += 1

            startx += 1
            starty += 1
            loop -= 1

        if mid == 1 :
            nums[int(n/2)][int(n/2)] = count

        return nums

print(Solution().generateMatrix(6))
# print(int(3/2))