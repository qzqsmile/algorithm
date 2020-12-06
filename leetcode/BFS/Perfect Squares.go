package BFS

//BFS，这种求最小步数，又不要求具体的访问路径的
//都可以往BFS上考虑
func numSquares(n int) int {
	var q []int
	q = append(q, n)
	step := 0
	visited := make(map[int] bool)
	visited[n] = true
	for len(q) > 0{
		var q1 []int
		for len(q) > 0{
			n1 := q[0]
			q = q[1:]
			if n1 == 0{
				return step
			}
			for i := 1; i*i <= n1; i++{
				if _, ok := visited[n1-(i*i)]; !ok{
					q1 = append(q1, n1-(i*i))
					visited[n1-(i*i)] = true
				}
			}
		}
		q = q1
		step++
	}
	return -1
}

//dp

func numSquares(n int) int {
	dp := []int{}
	for i := 0; i <= n; i++{
		dp = append(dp, int(^uint(0)>>3))
	}
	for i := 1; i*i <= n; i++{
		dp[i*i] = 1
	}

	for i := 1; i <= n; i++{
		for j := 1; j*j <= i; j++{
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[n]
}

func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}
