package main

/**
 * @param grid: a 2D integer grid
 * @return: an integer
 */
type coord struct {
	x int
	y int
}

func zombie(grid [][]int) int {
	// write your code here
	step := 0
	q := []coord{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				q = append(q, coord{i, j})
			}
		}
	}
	x := []int{1, 0, 0, -1}
	y := []int{0, -1, 1, 0}
	for len(q) > 0 {
		nq := q
		q = []coord{}
		for len(nq) > 0 {
			t := nq[0]
			nq = nq[1:]
			for i := 0; i < 4; i++ {
				tx := x[i] + t.x
				ty := y[i] + t.y

				if tx < 0 || ty < 0 || tx >= len(grid) || ty >= len(grid[0]) {
					continue
				}

				if grid[tx][ty] == 0 {
					grid[tx][ty] = 1
					q = append(q, coord{tx, ty})
				}
			}
		}
		if len(q) > 0 {
			step++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				return -1
			}
		}
	}
	return step
}

