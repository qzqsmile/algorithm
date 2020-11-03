class Solution:
    """
    @param: nums: A list of integers
    @return: A list of integers includes the index of the first number and the index of the last number
    """
    def subarraySumClosest(self, nums):
        # write your code here
        from collections import defaultdict
        mp = defaultdict(lambda:[])
        cur = 0
        tmp = []
        for i in range(0, len(nums)):
            cur += nums[i]
            mp[cur].append(i)
            tmp.append(cur)
            if cur == 0:
                return [0, i]
        tmp = sorted(tmp)
        b = [0, 0]
        closet = 100000
        for i in range(1, len(tmp)):
            if tmp[i] - tmp[i-1] < closet:
                closet = tmp[i] - tmp[i-1]
                b = sorted([mp[tmp[i]][0], mp[tmp[i-1]][0]])
                b[0] += 1
        return b
        
        
        