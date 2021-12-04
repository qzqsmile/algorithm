func rob(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    if len(nums) == 1{
        return nums[0]
    }
    ans1 := rob1(nums[0:len(nums)-1])
    ans2 := rob1(nums[1:len(nums)])
    return max(ans1, ans2)
}

func rob1(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    maxMoney := nums[0]
    dp := []int{0, nums[0]}
    for i := 1; i < len(nums); i++{
        dp = append(dp, max(nums[i]+dp[i-1], dp[i]))
        maxMoney = max(maxMoney, dp[len(dp)-1])
    }
    return maxMoney
}

func max(a int, b int)int{
    if  a > b{
        return a
    }
    return b
}