package unionfindset

func longestConsecutive(nums []int) int {
	maxLength := 0
	mp := make(map[int]bool)
	for _, v := range nums{
		mp[v] = true
	}

	for _, v := range nums{
		if _, ok := mp[v-1]; !ok {
			l, t := 0, v
			for;;{
				if _, ok := mp[t]; ok{
					l++
					t++
				}else{
					break
				}
			}
			maxLength = max(l, maxLength)
		}
	}

	return maxLength
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}
