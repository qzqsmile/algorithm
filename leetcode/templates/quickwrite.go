sort.Slice(intervals, func(i, j int)bool{
	if intervals[i][0] < intervals[j][0]{
		return true
	}
	return false
})