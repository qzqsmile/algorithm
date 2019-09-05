class Solution:
    def subsets(self, nums):
        res = [[]]
        if nums:
            self.backtracing(res, nums)
        return res

    def backtracing(self, res, nums):
        if len(nums) == 1:
            res.append(nums)
            return
        l = len(nums)
        self.backtracing(res, nums[0:l - 1])
        val = nums[-1]
        tmp = []
        print(res)
        import pdb; pdb.set_trace()
        for each in res:
            each.append(val)
            tmp.append(each)
        res += tmp

s = Solution()
res = s.subsets([1,2,3])
print(res)

