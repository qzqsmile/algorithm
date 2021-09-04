package dp

//坐标型
func uniquePaths(m int, n int) int {
	var last_row []int
	for i := 0; i < n; i++{
		last_row = append(last_row, 1)
	}
	cur_row := append([]int{}, last_row...)

	for i := 1; i < m; i++{
		for j := 0; j < n; j++{
			if j == 0{
				cur_row[j] = last_row[j]
			}else{
				cur_row[j] = last_row[j] + cur_row[j-1]
			}
		}
		last_row = append([]int{}, cur_row...)
	}
	return cur_row[len(cur_row)-1]
}


