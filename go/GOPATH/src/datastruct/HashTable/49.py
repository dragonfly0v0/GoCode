from typing import List

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        if len(strs) == 0 :
            return ""
        
        from collections import defaultdict
        # groups = defaultdict(list) # 初始化哈希表，key为排序后的字符串，value为字母异位词组成的切片
    
        # # 遍历字符串数组
        # for s in strs:
        #     # 将字符串排序后作为 key
        #     key = ''.join(sorted(s))
        #     print(key)
        #     # 将字符串添加到 key 对应的 value 列表中
        #     groups[key].append(s)
        
        # # 将哈希表中的 value 组成一个列表作为结果返回
        #     print(groups)

        groups = {} # 初始化哈希表，key为长度为26的全0列表，value为字母异位词组成的切片
        
        # 遍历字符串数组
        for s in strs:
            # 统计每个字符出现的次数
            cnt = [0] * 26
            for ch in s:
                cnt[ord(ch) - ord('a')] += 1
            
            # 将cnt转换为元组，作为 key
            key = tuple(cnt)
            
            # 将字符串添加到 key 对应的 value 列表中
            if key not in groups:
                groups[key] = []
            groups[key].append(s)

        print(groups)
        
        # 将哈希表中的 value 组成一个列表作为结果返回
        return list(groups.values())
        
strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
res = Solution().groupAnagrams(strs)