package dp

func rob(nums []int) int {
	if len(nums) == 0{
		return 0
	}else if len(nums) == 1{
		return nums[0]
	}
	dp1, dp2, dp3 := 0, nums[0], nums[1]
	for i := 2; i < len(nums); i++{
		ndp := max(dp1, dp2)+nums[i]
		dp1, dp2, dp3 = dp2, dp3, ndp
	}
	return max(dp2, dp3)
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}


