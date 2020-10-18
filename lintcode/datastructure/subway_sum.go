package main

/**
 * @param nums: A list of integers
 * @return: A list of integers includes the index of the first number and the index of the last number
 */
func subarraySum (nums []int) []int {
	// write your code here
	mp := make(map[int]int)
	preSum := 0
	mp[0] = 0

	for i := 0; i < len(nums); i++{
		preSum += nums[i]
		if v, ok := mp[preSum]; ok{
			return []int{v, i}
		}
		mp[preSum] = i + 1
	}
	return []int{0, 0}
}

