class Solution:
    def maximumGap(self, nums) -> int:
        num = self.radixSort(nums)
        res = 0
        for i in range(1, len(num)):
            res = max(num[i] - num[i - 1], res)
        return res

    def radixSort(self, nums):
        for i in range(0, 32):
            needle = 1 << i
            zero = []
            one = []
            import pdb;
            pdb.set_trace()
            for each in nums:
                if each & needle != 0:
                    one.append(each)
                else:
                    zero.append(each)
            nums = []
            nums += zero
            nums += one
        return nums

s = Solution()
res = s.maximumGap([3, 6, 9, 1])

