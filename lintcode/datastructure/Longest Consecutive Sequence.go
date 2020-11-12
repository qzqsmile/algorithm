package main

func longestConsecutive(nums []int) int {
	maxLength := 0
	mp := make(map[int] int)
	for _, v := range nums{
		mp[v] = 1
	}
	for _, v := range nums{
		if _, ok := mp[v]; ok{
			length := 1
			t := v
			for {
				if _, ok := mp[t-1]; ok{
					length += mp[t-1]
					delete(mp, t-1)
					t--
				}else{
					break
				}
			}
			mp[v] = max(length, mp[v])
			maxLength = max(length, maxLength)
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