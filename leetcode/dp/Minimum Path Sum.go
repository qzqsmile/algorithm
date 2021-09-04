package dp

//坐标性dp，一般有个坐标。走到某个地方
func minPathSum(grid [][]int) int {
	if len(grid) == 0{
		return 0
	}
	last_row := append([]int{}, grid[0]...)
	for i := 1; i < len(last_row); i++{
		last_row[i] += last_row[i-1]
	}
	current_row := append([]int{}, last_row...)
	for i := 1; i < len(grid); i++{
		for j := 0; j < len(grid[0]); j++{
			if j == 0{
				current_row[j] = last_row[j] + grid[i][j]
			}else{
				current_row[j] = min(last_row[j], current_row[j-1]) + grid[i][j]
			}
		}
		last_row = append([]int{}, current_row...)
	}

	return current_row[len(current_row)-1]
}

func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}
