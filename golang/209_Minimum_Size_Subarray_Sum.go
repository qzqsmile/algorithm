package main

func minSubArrayLen(s int, nums []int) int {
	min := 0
	tmp_sum := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		tmp_sum += nums[j]
		for ; tmp_sum >= s; {
			if (min >= j-i+1) || (min == 0) {
				min = j - i + 1
			}
			tmp_sum -= nums[i]
			i++
		}
	}
	return min
}
