package main

func minSubArrayLen(s int, nums []int) int {
	minLength := len(nums) + 1
	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for ; i <= j && sum >= s; i++ {
			if j-i+1 < minLength {
				minLength = j - i + 1
			}
			sum -= nums[i]
		}
	}
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}
