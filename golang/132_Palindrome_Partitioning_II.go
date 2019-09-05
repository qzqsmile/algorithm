func minCut(s string) int {
	n := len(s)
	var dp = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i - 1
	}
	for i := 0; i < n; i++ {

		for j := 0; (i+j < n) && (i-j >= 0) && (s[i+j] == s[i-j]); j++ {
			dp[i+j+1] = min(dp[i+j+1], dp[i-j]+1)
		}
		for j := 0; (i+j < n) && (i-j-1 >= 0) && (s[i+j] == s[i-j-1]); j++ {
			dp[i+j+1] = min(dp[i+j+1], dp[i-j-1]+1)
		}
	}
	return dp[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
