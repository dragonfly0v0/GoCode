from typing import List

list = [10, 90, 170, 210 , 430, 876]

def BinarySerach(nums: List[int], target: int):
    left = 0
    right = len(nums) - 1

    while left <= right:
        mid = int(left + (right - left)/2)
        if nums[mid] == target:
            print("target {} 在列表中的位置是{}".format(target, mid))
            return
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

BinarySerach(list, 210)
