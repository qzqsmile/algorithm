class Solution:
    def search(self, nums, target):
        b, e = 0, len(nums)-1
        while b <= e:
            m = (b+e) // 2
            if nums[m] == target:
                return m
            elif nums[m] < target:
                if nums[m] >= nums[b] or nums[b] > target:
                    b = m+1
                else:
                    e = m-1
            else:
                if nums[m] >= nums[b] and nums[b] > target:
                    b = m+1
                else:
                    e = m-1
        return -1