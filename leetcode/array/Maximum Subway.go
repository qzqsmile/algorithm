package array

import "math"

func maxSubArray(nums []int) int {
	maxNum := math.MinInt32
	minNum := 0
	var sumArr []int
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		sumArr = append(sumArr, s)
	}

	for i := 0; i < len(sumArr); i++ {
		maxNum = max(maxNum, sumArr[i]-minNum)
		minNum = min(minNum, sumArr[i])
	}
	return maxNum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

