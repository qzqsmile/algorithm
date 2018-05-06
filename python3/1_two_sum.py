class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        mapping = {}
        for index, value in enumerate(nums):
            if value in mapping:
                return [mapping[value], index]
            else:
                mapping[target-value] = index
