//动规，最好分配m+1, n+1多分配一行和一列。可以少考虑很多问题。

func findLength(nums1 []int, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    dp := make([][]int, m+1)
    for i := 0; i < len(dp); i++{
        dp[i] = make([]int, n+1)
    }
    for i := 1; i < len(dp); i++{
        for j := 1; j < len(dp[0]); j++{
            if nums1[i-1] == nums2[j-1]{
                dp[i][j] = dp[i-1][j-1]+1
            }else{
                dp[i][j] = 0
            }
        }
    }
    ans := 0
    for i := 0; i < len(dp); i++{
        for j := 0; j < len(dp[0]); j++{
            if dp[i][j] > ans{
                ans = dp[i][j]
            }
        }
    }
    return ans
}