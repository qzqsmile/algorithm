//对于string类型的dp
//dp的设置的下标，最好与string相同，这样不容易出错

func longestValidParentheses(s string) int {

	maxLength := 0
	dp := []int{0}
	for i := 1; i < len(s); i++ {
		tmp := 0
		if s[i] == ')' {
			last := i - dp[i-1] - 1
			if last >= 0 && s[last] == '(' {
				if last-1 >= 0 {
					tmp = 2 + dp[i-1] + dp[last-1]
				} else {
					tmp = 2 + dp[i-1]
				}
			}
		}
		dp = append(dp, tmp)
	}

	for i := 0; i < len(dp); i++ {
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(str string) bool {
	if len(str) == 0 {
		return true
	}
	stk := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != ')' {
			stk = append(stk, str[i])
		} else {
			if len(stk) == 0 || stk[len(stk)-1] != '(' {
				return false
			}
			stk = stk[0 : len(stk)-1]
		}
	}
	return !(len(stk) > 0)
}
