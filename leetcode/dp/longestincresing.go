package dp

//dp version
func lengthOfLIS(nums []int) int {
    dp := []int{1}
    mlis:= 1
    for i := 1; i < len(nums); i++{
        m := 1
        for j := 0; j < i; j++{
            if nums[i] > nums[j] && dp[j]+1 > m{
                m = dp[j]+1
            }
        }
        dp = append(dp, m)
    }
    for _, v := range dp{
        if v > mlis{
            mlis = v
        }
    }
    return mlis
}