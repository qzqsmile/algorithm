func longestCommonSubsequence(text1 string, text2 string) int {
    r, c := len(text1), len(text2)
    dp := make([][]int, r+1)

    for i := 0; i < len(dp); i++{
        dp[i] = make([]int, c+1)
    }

    for i := 1; i < len(dp); i++{
        for j := 1; j < len(dp[0]); j++{
            if text1[i-1] == text2[j-1]{
                dp[i][j] = dp[i-1][j-1] + 1
            }else{
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    return dp[len(dp)-1][len(dp[0])-1]
}

func max(a int, b int)int{
    if a > b{
        return a
    }
    return b
}