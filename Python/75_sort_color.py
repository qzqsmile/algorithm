class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        b, e, i = 0,len(nums)-1, 0
        while i <= e:
            while nums[i] == 2 and i < e:
                nums[i], nums[e] = nums[e], nums[i]
                e -= 1
            while nums[i] == 0 and i > b:
                nums[i], nums[b] = nums[b], nums[i]
                b += 1
            i += 1