class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = [[]]
        from collections import defaultdict
        nums = sorted(nums)
        mp = defaultdict(lambda: 0)
        for each in nums:
            mp[each] += 1
        if nums:
            self.backtracing(res, nums, mp)
        return res

    def backtracing(self, res, nums, mp):
        if len(nums) == 1:
            res.append(nums)
            return
        l = len(nums)
        self.backtracing(res, nums[0:l - 1], mp)
        val = nums[-1]
        import copy
        tmp = copy.deepcopy(res)
        count = len(nums) - nums.index(val)
        for each in tmp:
            tmp = 0 if val not in each else len(each) - each.index(val)
            if mp[val] == 1 or count - 1 == tmp:
                each.append(val)
                res.append(each)