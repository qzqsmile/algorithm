class Solution:
    def findSubsequences(self, nums: List[int]) -> List[List[int]]:
        res, tmp = [], []
        self.subset(nums, 0, tmp, res)
        return res

    def subset(self, nums, index, tmp, res):
        if index >= len(nums):
            return
        used = {}
        for i in range(index, len(nums)):
            if nums[i] not in used and (not tmp or nums[i] >= tmp[-1]):
                used[nums[i]] = True
                tmp.append(nums[i])
                if len(tmp) >= 2:
                    res.append(tmp[:])
                self.subset(nums, i + 1, tmp, res)
                tmp.pop()



