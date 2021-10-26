package dp

func longestPalindrome(s string) string {
	longestStr := ""
	for i := 0; i < len(s); i++ {
		j, k := i, i
		for j >= 0 && k < len(s) {
			if s[j] == s[k] {
				longestStr = max(longestStr, s[j:k+1])
				j--
				k++
			} else {
				break
			}
		}

		j, k = i, i+1
		for j >= 0 && k < len(s) {
			if s[j] == s[k] {
				longestStr = max(longestStr, s[j:k+1])
				j--
				k++
			} else {
				break
			}
		}
	}
	return longestStr
}

func max(a string, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}