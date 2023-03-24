package hashtable

func leastBricks(wall [][]int) int {
	mp := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			mp[sum]++
		}
	}

	res := len(wall)
	for _, v := range mp {
		res = min(res, len(wall)-v)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
