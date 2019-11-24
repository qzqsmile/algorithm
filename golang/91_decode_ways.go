package main

import (
	"strconv"
)

func numDecodings(s string) int {
	dp := []int{1, 0}
	if len(s) != 0 && s[0] != '0' {
		dp[1] = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '0' && (s[i-1] == '0' || !isCanConvert(s[i-1:i+1])) {
			dp = append(dp, 0)
		} else if s[i-1] == '0' || !isCanConvert(s[i-1:i+1]) {
			dp = append(dp, dp[i])
		} else if s[i] == '0' {
			dp = append(dp, dp[i-1])
		} else {
			dp = append(dp, dp[i]+dp[i-1])
		}
	}
	return dp[len(dp)-1]
}

func isCanConvert(s string) bool {
	i, _ := strconv.Atoi(s)
	if i <= 26 {
		return true
	} else {
		return false
	}
}
