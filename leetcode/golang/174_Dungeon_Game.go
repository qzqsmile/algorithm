]
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	mat := make([][]int, m+1)
	intMax := 1<<63 - 1
	for i := range mat {
		mat[i] = make([]int, n+1)
		for j := range mat[i] {
			mat[i][j] = intMax
		}
	}
	mat[m][n-1], mat[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mat[i][j] = max(min(mat[i+1][j], mat[i][j+1])-dungeon[i][j], 1)
		}
	}
	return mat[0][0]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
