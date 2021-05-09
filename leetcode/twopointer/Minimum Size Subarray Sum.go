package twopointer

import "math"

func minSubArrayLen(s int, nums []int) int {
	left, right := 0, 0
	sum, min := 0, math.MaxInt64
	for right < len(nums){
		c := nums[right]
		right++
		sum += c
		for sum >= s{
			if right-left < min{
				min = right-left
			}
			c := nums[left]
			sum -= c
			left++
		}
	}
	if min == math.MaxInt64{
		return 0
	}
	return min
}