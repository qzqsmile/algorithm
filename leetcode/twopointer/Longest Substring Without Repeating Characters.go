package twopointer

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0{
		return 0
	}
	i, j := 0, 0
	maxLongest := 1
	mp := make(map[uint8]int)
	for ;i < len(s);{
		for;j < len(s);j++{
			if _, ok := mp[s[j]];ok && mp[s[j]] >= i{
				break
			}
			mp[s[j]] = j
		}
		if j - i > maxLongest{
			maxLongest = j - i
		}
		if j < len(s){
			i = mp[s[j]]+1
			mp[s[j]] = j
			j++
		}else{
			break
		}
	}
	return maxLongest
}


