package dp

func canJump(nums []int) bool {
	if len(nums) <= 1{
		return true
	}
	dp := append([]int{}, nums[0])
	for i := 1; i < len(nums); i++{
		if dp[i-1] >= i{
			dp = append(dp, max(dp[i-1], i+nums[i]))
		}else{
			return false
		}
	}
	return true
}

// func max(a int, b int) int{
// 	if a > b{
// 		return a
// 	}
// 	return b
// }
