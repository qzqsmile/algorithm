class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        from collections import defaultdict
        mp, prefix_sum = defaultdict(lambda: []), 0
        mp[0].append(0)
        count = 0
        for i, n in enumerate(nums):
            prefix_sum += n
            if prefix_sum-k in mp:
                count += len(mp[prefix_sum-k])
            mp[prefix_sum].append(i+1)
        return count
