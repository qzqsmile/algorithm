package backtrack

//超时
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) != 0 && len(p) == 0 {
		return false
	} else if len(s) == 0 && len(p) != 0 {
		if len(p) >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		}
		return false
	} else {
		if s[0] == p[0] || p[0] == '.' {
			if len(p) >= 2 && p[1] == '*' {
				return isMatch(s, p[2:]) || isMatch(s[1:], p[2:]) || isMatch(s[1:], p)
			}
			return isMatch(s[1:], p[1:])
		} else if len(p) >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		} else {
			return false
		}
	}
}

//recuresion
