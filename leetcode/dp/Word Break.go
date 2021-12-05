func wordBreak(s string, wordDict []string) bool {
    mp := make(map[string] bool)
    dp := make([]bool, len(s)+1)
    dp[0] = true

    for _, v := range wordDict{
        mp[v] = true
    }

    for i := 0; i < len(s); i++{
        for k, _ := range mp{
            if i+1-len(k) >= 0 && dp[i+1-len(k)] 
            && s[i+1-len(k): i+1] == k{
                dp[i+1] = true 
                break
            }
        }
    }
    fmt.Println(dp)
    return dp[len(dp)-1]
}