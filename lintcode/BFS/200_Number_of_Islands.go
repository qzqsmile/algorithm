package main

type coord1 struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				bfs(grid, i, j)
			}
		}
	}
	return ans
}

func bfs(grid [][]byte, i int, j int) {
	x := []int{0, 1, -1, 0}
	y := []int{1, 0, 0, -1}
	q := []coord{}
	q = append(q, coord{i, j})
	grid[i][j] = '0'

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			xt := t.x + x[i]
			yt := t.y + y[i]

			if xt < 0 || yt < 0 || xt >= len(grid) || yt >= len(grid[0]) {
				continue
			}
			if grid[xt][yt] == '1' {
				grid[xt][yt] = '0'
				q = append(q, coord{xt, yt})
			}
		}
	}
}


