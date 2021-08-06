class Solution:
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        S = {-1: []}
        import copy
        for x in sorted(nums):
            S[x] = copy.copy(max((S[d] for d in S if x % d == 0), key=len)) +[x]
        return max(S.values(), key=len)
