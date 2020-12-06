package BFS

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[0]); j++{
			if grid[i][j] == '1'{
				bfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

type corridate struct {
	x int
	y int
}

// 这种错误的写法根源在于会重复入队列。对于Tree这种结构，其实
// 天然不用考虑重复入队的问题。这里会重复入队，虽然对结果没有影响
// 但是会导致time limited的问题。解决方法就是入队时就对其进行操作
// 防止其重复入队
// 另一个值得学习的点在于数组表示坐标移动
func bfs(grid [][]byte, i int, j int){
	q := [] corridate{}
	q = append(q, corridate{i, j})
	// grid[i][j] = '0'
	for len(q) > 0{
		n := q[0]
		q = q[1:]
		grid[n.x][n.y] = '0'
		if n.x-1>=0 && grid[n.x-1][n.y] == '1'{
			q = append(q, corridate{n.x-1, n.y})
		}

		if n.x+1<len(grid) && grid[n.x+1][n.y] == '1'{
			q = append(q, corridate{n.x+1, n.y})
		}

		if n.y-1>=0 && grid[n.x][n.y-1] == '1'{
			q = append(q, corridate{n.x, n.y-1})
		}

		if n.y+1<len(grid[0]) && grid[n.x][n.y+1] == '1'{
			q = append(q, corridate{n.x, n.y+1})
		}
	}
}

//right
func bfs(grid [][]byte, i int, j int){
	x := []int{0, 1, -1, 0}
	y := []int{1, 0, 0, -1}
	q := [] corridate{}
	q = append(q, corridate{i, j})
	grid[i][j] = '0'
	for len(q) > 0{
		n := q[0]
		q = q[1:]
		for i := 0; i < 4; i++{
			xt := n.x + x[i]
			yt := n.y + y[i]
			if xt < 0 || yt < 0 || xt >= len(grid) || yt >= len(grid[0]){
				continue
			}
			if grid[xt][yt] == '1'{
				grid[xt][yt] = '0'
				q = append(q, corridate{xt, yt})
			}
		}
	}
}
