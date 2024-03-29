func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)

    for i := 0; i < len(dp); i++{
        dp[i] = make([]int, n+1)
    }
    
    for i := 0; i < len(dp[0]); i++{
        dp[0][i] = i
    }

    for i := 0; i < len(dp); i++{
        dp[i][0] = i
    }
    for i := 1; i < len(dp); i++{
        for j := 1; j < len(dp[0]); j++{
            dp[i][j] =1+ min(dp[i-1][j], dp[i][j-1])
            if word1[i-1] == word2[j-1]{
                dp[i][j] = min(dp[i][j], dp[i-1][j-1])
            }else{
                dp[i][j] = min(dp[i][j], 1+dp[i-1][j-1])
            }
        }  
    }
    return dp[m][n]
}


func min(a int, b int) int{
    if a < b{
        return a
    }
    return b
}