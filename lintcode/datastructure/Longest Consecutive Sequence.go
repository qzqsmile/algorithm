package main

func longestConsecutive(nums []int) int {
	maxLength := 0
	mp := make(map[int] bool)
	for _, v := range nums{
		mp[v] = true
	}
	for _, v := range nums{
		if _, ok := mp[v-1]; !ok{
			length := 1
			for {
				if _, ok := mp[v+1]; ok{
					length++
					v++
				}else{
					break
				}
			}
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
