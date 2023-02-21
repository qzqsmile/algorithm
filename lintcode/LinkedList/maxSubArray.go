package main

/**
 * @param nums: A list of integers
 * @return: A integer indicate the sum of max subarray
 */
func maxSubArray(nums []int) int {
	// write your code here
	if len(nums) == 0 {
		return 0
	}
	pre, cur, res := 0, 0, nums[0]

	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		res = max(res, nums[i])
		if cur < pre {
			pre = cur
		} else {
			res = max(res, cur-pre)
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
